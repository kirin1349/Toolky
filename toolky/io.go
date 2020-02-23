package toolky

import (
	"path/filepath"
	"os"
	"strings"
	"strconv"
	"os/exec"
	"fmt"
	"io/ioutil"
)

func GetAppDir() (string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func HasOSArgByKey(key string) (bool) {
	for i := range os.Args {
		if os.Args[i] == key {
			return true
		}
	}
	return false
}

func GetOSArgByKey(key string) (string) {
	for i := range os.Args {
		if os.Args[i] == key {
			if len(os.Args) > i + 1 {
				if strings.Index(os.Args[i + 1], "-") == 0 {
					return ""
				}
				str := os.Args[i + 1]
				str = strings.Replace(str, "\"", "", -1)
				str = strings.Replace(str, "'", "", -1)
				return str
			}
		}
	}
	return ""
}

func GetOSArgNumByKey(key string) (int) {
	for i := range os.Args {
		if os.Args[i] == key {
			if len(os.Args) > i + 1 {
				str := os.Args[i + 1]
				str = strings.Replace(str, "\"", "", -1)
				str = strings.Replace(str, "'", "", -1)
				argNum, err := strconv.Atoi(str)
				if err == nil {
					return argNum
				}
			}
		}
	}
	return 0
}

func QuickExec(cmdStr string, printErr bool) (bool) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	err := cmd.Run()
	if err != nil {
		if printErr {
			fmt.Println(cmdStr + " failed with error " + err.Error())
		}
		return false
	}
	return true
}

func QuickExecAt(dir string, cmdStr string, printErr bool) (bool) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		if printErr {
			fmt.Println(cmdStr + " failed with error " + err.Error())
		}
		return false
	}
	return true
}

func QuickExecWithOutput(cmdStr string, printErr bool) (bool, string) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	opt, err := cmd.Output()
	if err != nil {
		if printErr {
			fmt.Println(cmdStr + " failed with error " + err.Error())
		}
		return false, ""
	}
	return true, string(opt)
}

func QuickExecWithOutputAt(dir string, cmdStr string, printErr bool) (bool, string) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Dir = dir
	opt, err := cmd.Output()
	if err != nil {
		if printErr {
			fmt.Println(cmdStr + " failed with error " + err.Error())
		}
		return false, ""
	}
	return true, string(opt)
}

func QuickExecWithStdOutput(cmdStr string, printErr bool) (bool) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		if printErr {
			fmt.Println(cmdStr + " failed with error " + err.Error())
		}
		return false
	}
	return true
}

func QuickExecWithStdOutputAt(dir string, cmdStr string, printErr bool) (bool) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if printErr {
			fmt.Println(cmdStr + " failed with error " + err.Error())
		}
		return false
	}
	return true
}

func QuickRead(filePath string, printErr bool) (bool, string) {
	buff, err := ioutil.ReadFile(filePath)
	if err != nil {
		if printErr {
			fmt.Println("open file " + filePath + " failed with error " + err.Error())
		}
		return false, ""
	}
	return true, string(buff)
}

func QuickWrite(filePath string, fileContent string, printErr bool) (bool) {
	err := ioutil.WriteFile(filePath, []byte(fileContent), os.ModeAppend)
	if err != nil {
		if printErr {
			fmt.Println("write file " + filePath + " failed with error " + err.Error())
		}
		return false
	}
	return true
}

func BuildCMDCopy(srcFile string, destFile string, bRecursion bool, bSudo bool) (string) {
	cmdStr := ""
	if bSudo {
		cmdStr += "sudo "
	}
	cmdStr += "cp "
	if bRecursion {
		cmdStr += "-R "
	}
	cmdStr += srcFile + " " + destFile
	return cmdStr
}

func BuildCMDRemove(destFile string, bRecursion bool, bSudo bool) (string) {
	cmdStr := ""
	if bSudo {
		cmdStr += "sudo "
	}
	cmdStr += "rm "
	if bRecursion {
		cmdStr += "-rf "
	} else {
		cmdStr += "-f "
	}
	cmdStr += destFile
	return cmdStr
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFolder(path string) (error) {
	result, err := PathExists(path)
	if err != nil {
		return err
	}
	if result == false {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFile(path string) (error) {
	result, err := PathExists(path)
	if err != nil {
		return err
	}
	if result == false {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func RemoveFile(path string) (error) {
	result, err := PathExists(path)
	if err != nil {
		return err
	}
	if result {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func CombinePath(folder string, file string) (string) {
	if folder == "" {
		return file
	}
	folder = strings.Replace(folder, "\\", "/", -1)
	index := strings.LastIndex(folder, "/")
	if index != (len(folder) - 1) {
		folder += "/"
	}
	return folder + file
}