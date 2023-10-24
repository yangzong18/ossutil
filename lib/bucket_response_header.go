package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var specChineseResponseHeader = SpecText{
	synopsisText: "bucket的响应头设置",

	paramText: "bucket_url [local_xml_file] [options]",

	syntaxText: ` 
    ossutil response-header --method put oss://bucket local_xml_file [options]
    ossutil response-header --method get oss://bucket [local_xml_file] [options]
    ossutil response-header --method delete oss://bucket [options]
`,
	detailHelpText: `
    response-header 命令通过设置method选项值为put、get、delete,可以设置、查询或者删除bucket的响应头设置;

用法:
    该命令有三种用法:

    1) ossutil response-header --method put oss://bucket local_xml_file [options]
        这个命令从配置文件local_xml_file中读取响应头设置,然后设置bucket的响应头设置,
        配置文件是一个xml格式的文件,举例如下
   
        <?xml version="1.0" encoding="UTF-8"?>
        <ResponseHeaderConfiguration>
			<Rule>
				<Name>rule1</Name>
				<Filters>
					<Operation>Put</Operation>
					<Operation>GetObject</Operation>
				</Filters>
				<HideHeaders>
					<Header>Last-Modified</Header>
					<Header>x-oss-request-id</Header>
				</HideHeaders>
			</Rule>
			<Rule>
				<Name>rule2</Name>
				<Filters>
					<Operation>*</Operation>
				</Filters>
				<HideHeaders>
					<Header>Last-Modified</Header>
					<Header>x-oss-request-id</Header>
				</HideHeaders>
			</Rule>
		</ResponseHeaderConfiguration>

    2) ossutil response-header --method get oss://bucket [local_xml_file] [options]
        这个命令查询bucket的响应头设置
        如果输入参数local_xml_file，响应头设置将输出到该文件，否则输出到屏幕上

    3) ossutil response-header --method delete oss://bucket [options]
        这个命令删除bucket的响应头设置
`,

	sampleText: `
    1) 设置bucket的响应头设置
       ossutil response-header --method put oss://bucket local_xml_file

    2) 查询bucket的响应头设置，结果输出到标准输出
       ossutil response-header --method get oss://bucket

	3) 查询bucket的响应头设置，结果输出到本地文件
       ossutil response-header --method get oss://bucket local_xml_file
	
    4) 删除bucket的响应头设置
       ossutil response-header --method delete oss://bucket
`,
}

var specEnglishResponseHeader = SpecText{
	synopsisText: "manage bucket's response header configuration",

	paramText: "bucket_url [local_xml_file] [options]",

	syntaxText: ` 
    ossutil response-header --method put oss://bucket local_xml_file [options]
    ossutil response-header --method get oss://bucket [local_xml_file] [options]
    ossutil response-header --method delete oss://bucket [options]
`,
	detailHelpText: ` 
    response-header command can set, get and delete the response header configuration of 
    the oss bucket by setting method option value to put, get and delete

Usage:
    There are three usages for this command:
	
    1) ossutil response-header --method put oss://bucket local_xml_file [options]
        The command sets the response header configuration of bucket from local file local_xml_file
        The local_xml_file is xml format,for example

        <?xml version="1.0" encoding="UTF-8"?>
        <ResponseHeaderConfiguration>
			<Rule>
				<Name>rule1</Name>
				<Filters>
					<Operation>Put</Operation>
					<Operation>GetObject</Operation>
				</Filters>
				<HideHeaders>
					<Header>Last-Modified</Header>
					<Header>x-oss-request-id</Header>
				</HideHeaders>
			</Rule>
			<Rule>
				<Name>rule2</Name>
				<Filters>
					<Operation>*</Operation>
				</Filters>
				<HideHeaders>
					<Header>Last-Modified</Header>
					<Header>x-oss-request-id</Header>
				</HideHeaders>
			</Rule>
		</ResponseHeaderConfiguration>
        
    2) ossutil response-header --method get oss://bucket [local_xml_file] [options]
       The command gets the response header configuration of bucket
       if you input parameter local_xml_file,the configuration will be output to local_xml_file
       if you don't input parameter local_xml_file,the configuration will be output to stdout	

    3) ossutil response-header --method delete oss://bucket ruleID
       The command delete the response header configuration of bucket
`,
	sampleText: ` 
    1) put bucket response header configuration
       ossutil response-header --method put oss://bucket local_xml_file

    2) get bucket response header configuration, output to stdout
       ossutil response-header --method get oss://bucket

	3) get bucket response header configuration, output to local file
       ossutil response-header --method get oss://bucket local_xml_file
	
    4) delete bucket response header configuration
       ossutil response-header --method delete oss://bucket
`,
}

type ResponseHeaderCommand struct {
	command    Command
	bucketName string
}

var responseHeaderCommand = ResponseHeaderCommand{
	command: Command{
		name:        "response-header",
		nameAlias:   []string{"response-header"},
		minArgc:     1,
		maxArgc:     2,
		specChinese: specChineseResponseHeader,
		specEnglish: specEnglishResponseHeader,
		group:       GroupTypeNormalCommand,
		validOptionNames: []string{
			OptionConfigFile,
			OptionEndpoint,
			OptionAccessKeyID,
			OptionAccessKeySecret,
			OptionSTSToken,
			OptionProxyHost,
			OptionProxyUser,
			OptionProxyPwd,
			OptionLogLevel,
			OptionMode,
			OptionECSRoleName,
			OptionTokenTimeout,
			OptionRamRoleArn,
			OptionRoleSessionName,
			OptionReadTimeout,
			OptionConnectTimeout,
			OptionSTSRegion,
			OptionMethod,
			OptionItem,
			OptionSkipVerifyCert,
			OptionUserAgent,
			OptionSignVersion,
			OptionRegion,
			OptionCloudBoxID,
			OptionForcePathStyle,
		},
	},
}

// function for FormatHelper interface
func (rhc *ResponseHeaderCommand) formatHelpForWhole() string {
	return rhc.command.formatHelpForWhole()
}

func (rhc *ResponseHeaderCommand) formatIndependHelp() string {
	return rhc.command.formatIndependHelp()
}

// Init simulate inheritance, and polymorphism
func (rhc *ResponseHeaderCommand) Init(args []string, options OptionMapType) error {
	return rhc.command.Init(args, options, rhc)
}

// RunCommand simulate inheritance, and polymorphism
func (rhc *ResponseHeaderCommand) RunCommand() error {
	strMethod, _ := GetString(OptionMethod, rhc.command.options)

	if strMethod == "" {
		return fmt.Errorf("--method value is empty")
	}

	strMethod = strings.ToLower(strMethod)
	if strMethod != "put" && strMethod != "get" && strMethod != "delete" {
		return fmt.Errorf("--method value is not in the optional value:put|get|delete")
	}

	srcBucketUrL, err := GetCloudUrl(rhc.command.args[0], "")
	if err != nil {
		return err
	}

	rhc.bucketName = srcBucketUrL.bucket
	switch strMethod {
	case "put":
		err = rhc.PutBucketReplication()
	case "get":
		err = rhc.GetBucketResponseHeader()
	case "delete":
		err = rhc.DeleteBucketResponseHeader()
	}
	return err
}

func (rhc *ResponseHeaderCommand) PutBucketReplication() error {
	if len(rhc.command.args) < 2 {
		return fmt.Errorf("put bucket response header need at least 2 parameters,the local xml file is empty")
	}

	xmlFile := rhc.command.args[1]
	fileInfo, err := os.Stat(xmlFile)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("%s is dir,not the expected file", xmlFile)
	}

	if fileInfo.Size() == 0 {
		return fmt.Errorf("%s is empty file", xmlFile)
	}

	// parsing the xml file
	file, err := os.Open(xmlFile)
	if err != nil {
		return err
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	client, err := rhc.command.ossClient(rhc.bucketName)
	if err != nil {
		return err
	}
	return client.PutBucketResponseHeaderXml(rhc.bucketName, string(text))
}

func (rhc *ResponseHeaderCommand) confirm(str string) bool {
	var val string
	fmt.Printf(getClearStr(fmt.Sprintf("cors: overwrite \"%s\"(y or N)? ", str)))
	if _, err := fmt.Scanln(&val); err != nil || (strings.ToLower(val) != "yes" && strings.ToLower(val) != "y") {
		return false
	}
	return true
}

func (rhc *ResponseHeaderCommand) GetBucketResponseHeader() error {
	client, err := rhc.command.ossClient(rhc.bucketName)
	if err != nil {
		return err
	}
	output, err := client.GetBucketResponseHeaderXml(rhc.bucketName)
	var outFile *os.File
	if len(rhc.command.args) >= 2 {
		fileName := rhc.command.args[1]
		_, err = os.Stat(fileName)
		if err == nil {
			bContinue := rhc.confirm(fileName)
			if !bContinue {
				return nil
			}
		}

		outFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0660)
		if err != nil {
			return err
		}
		defer outFile.Close()
	} else {
		outFile = os.Stdout
	}

	outFile.Write([]byte(output))

	fmt.Printf("\n\n")
	return nil
}

func (rhc *ResponseHeaderCommand) DeleteBucketResponseHeader() error {
	client, err := rhc.command.ossClient(rhc.bucketName)
	if err != nil {
		return err
	}
	return client.DeleteBucketResponseHeader(rhc.bucketName)
}
