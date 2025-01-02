package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"recruit/ent"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var sftpConfig *ssh.ClientConfig

func init() {
	// Initialize SFTP client configuration
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sftpConfig = &ssh.ClientConfig{
		User: os.Getenv("SFTP_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SFTP_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type FileHandler struct{}

func (FileHandler) Download(c *gin.Context) {
	var reqBody struct {
		SftpFilePath string `json:"filepath" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	client, err := connectSFTP()
	if err != nil {
		log.Println("Error connecting to SFTP")
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println("Error connecting to SFTP")
		return
	}
	defer client.Close()

	file, err := client.Open(reqBody.SftpFilePath)
	if err != nil {
		log.Println("Error Opening the file")
		c.JSON(http.StatusInternalServerError, gin.H{"Err": err.Error()})
		log.Println("Error Opening the file")
		return
	}
	defer file.Close()

	// Set the Content-Disposition header to prompt download
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", reqBody.SftpFilePath))

	// Print the file content as it is being read
	// fmt.Println("Reading file:")
	// content, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Println(string(content))

	// Copy the file content to the response writer
	fmt.Println("Reading file")
	fmt.Println(file)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		log.Println("Error copying file to response:", err)
		return
	}

}

func (FileHandler) PDFDownload(c *gin.Context) {
	var reqBody struct {
		SftpFilePath string `json:"filepath" binding:"required"`
	}

	//reqBody.SftpFilePath = "Pdf/ValueTest.pdf"
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	client, err := connectSFTP()
	if err != nil {
		log.Println("Error connecting to SFTP")
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println("Error connecting to SFTP")
		return
	}
	defer client.Close()
	log.Println(reqBody.SftpFilePath)
	file, err := client.Open(reqBody.SftpFilePath)

	log.Println(file)
	if err != nil {
		fmt.Println("error Opening the file ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Err": err.Error()})
		log.Println("Error Opening the file")
		return
	}
	defer file.Close()

	// Set the Content-Disposition header to prompt download
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", reqBody.SftpFilePath))

	// Print the file content as it is being read
	// fmt.Println("Reading file:")

	// content, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Println(string(content))

	// Copy the file content to the response writer
	log.Println("Reading file")
	log.Println(file)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		log.Println("Error copying file to response:", err)
		return
	}

}

func (FileHandler) Upload(c *gin.Context, dbclient *ent.Client) {

	envmode := os.Getenv("ENV_MODE")
	var sftpfolderpath string
	if envmode == "production" {
		sftpfolderpath = "/DeptExams/uploads/PICORSIGN/"
	} else {
		sftpfolderpath = "/DeptExams/uploads/uat/PICORSIGN/"
	}

	kyctype := c.PostForm("kyctype")
	doctype := c.PostForm("doctype")
	articleid := c.PostForm("articleid")
	//employeeid := c.PostForm("employeeid")
	name := c.PostForm("filename")
	folder := c.PostForm("folder")
	fname := c.PostForm("pdfname")
	examcode := c.PostForm("examcode")
	year := c.PostForm("year")
	if folder == "Pdf" {
		if examcode == "" || year == "" {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Enter all Fields values"})
			return
		}
	}

	fmt.Println(name + " " + folder)
	//document := c.PostForm("document")

	fmt.Println("Request Body:", c.Request.Form.Encode())
	file, err := c.FormFile("document")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "No file found in request"})
		return
	}

	if err := validateFile(file, kyctype, articleid, "TEST", doctype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	subfolder := articleid
	if kyctype == "PICORSIGN" {
		subfolder = folder
	}

	// timestamp := time.Now().Format("20060102150405.000")
	//filename := fmt.Sprintf("%s_%s.jpg", subfolder, doctype)
	filename := fmt.Sprintf("%s.jpg", name)
	fmt.Println(doctype + "--------------------")
	if doctype == "PDF" {
		filename = fmt.Sprintf("%s.pdf", name)
	}

	client, err := connectSFTP()
	if err != nil {
		log.Println("Error connecting to SFTP")
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println("Error connecting to SFTP")
		return
	}
	defer client.Close()
	// rootdir := os.Getenv("SFTP_ROOT_FOLDER")
	// fullpath := fmt.Sprintf("/%s/%s", kyctype, subfolder)
	dir := fmt.Sprintf("/DeptExams/uploads/%s/%s", kyctype, subfolder)
	// dir := filepath.Join("/", rootdir, fullpath)
	// err = createSFTPDirectory(client, dir)
	// if err != nil {
	// 	log.Println("Error creating SFTP directory")
	// 	c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	// 	log.Println("Error creating SFTP directory")
	// 	return
	// }
	//dir = "/DeptExams/uploads/PICORSIGN/" + folder
	dir = sftpfolderpath + folder

	fmt.Println(filename + "--------")
	destFile, err := client.Create(fmt.Sprintf("%s/%s", dir, filename))
	if err != nil {
		log.Println("Error creating File in the directory")
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println("Error creating File in the directory")
		return
	}
	defer destFile.Close()

	srcFile, err := file.Open()
	if err != nil {
		log.Println("Error opening File in the directory")
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println("Error opening File in the directory")
		return
	}
	defer srcFile.Close()

	_, err = destFile.ReadFrom(srcFile)
	if err != nil {
		log.Println("Error Reading File")
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println("Error Reading File")
		return
	}
	year1, _ := strconv.Atoi(year)
	code1, _ := strconv.Atoi(examcode)
	if doctype == "PDF" {

		_, err := dbclient.PDF.Create().SetPath(filepath.Join(dir, filename)).SetFilename(fname).SetYear(year1).SetExamcode(code1).Save(context.Background())

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Failed to Upload",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "File uploaded successfully",
		"filename": filename,
		"path":     filepath.Join(dir, filename),
	})
}

func contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

type ValidationError struct {
	Field string  `json:"field"`
	Msg   string  `json:"msg"`
	Value *string `json:"value,omitempty"`
}

func validateFile(file *multipart.FileHeader, kyctype, articleid, employeeid, doctype string) []ValidationError {
	errors := []ValidationError{}
	kycDocTypes := []string{"PIC", "SIGN", "PDF"}
	articleDocTypes := []string{
		"OTHR",
	}

	fmt.Println(file.Size, "File Size", doctype)

	if file.Size > 100000 && doctype == "PIC" {
		errors = append(errors, ValidationError{"document", "File size should be below 100 KB!!", nil})
	} else if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/jpg" && file.Header.Get("Content-Type") != "application/pdf" {
		// Throw an error because the file type is neither jpg nor jpeg
		errors = append(errors, ValidationError{"document", "only jpg/jpeg files are allowed", nil})
	}

	if kyctype == "PICORSIGN" {
		if doctype == "" {
			errors = append(errors, ValidationError{"doctype", "field is required", nil})
		} else if !contains(kycDocTypes, doctype) {
			errors = append(errors, ValidationError{"doctype", "invalid field value", &doctype})
		}
	} else if kyctype == "ArticleFiles" {
		if articleid == "" {
			errors = append(errors, ValidationError{"articleid", "field is required", nil})
		}
		if doctype == "" {
			errors = append(errors, ValidationError{"doctype", "field is required", nil})
		} else if !contains(articleDocTypes, doctype) {
			errors = append(errors, ValidationError{"doctype", "invalid field value", &doctype})
		}
	} else if kyctype == "" {
		errors = append(errors, ValidationError{"kyctype", "field is required", nil})
	} else {
		errors = append(errors, ValidationError{"kyctype", "invalid field value", &kyctype})
	}

	if len(errors) == 0 {
		return nil
	}

	return errors

}

func connectSFTP() (*sftp.Client, error) {
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("SFTP_URL"), os.Getenv("SFTP_PORT")), sftpConfig)
	if err != nil {
		return nil, err
	}
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		sshClient.Close()
		return nil, err
	}
	return sftpClient, nil
}

func createSFTPDirectory(client *sftp.Client, dir string) error {
	_, err := client.Stat(dir)
	if err == nil {
		// Directory already exists
		return nil
	}
	err = client.Mkdir(dir)
	if err != nil {
		return err
	}
	return nil
}
