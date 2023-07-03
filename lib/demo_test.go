package lib

import (
	. "gopkg.in/check.v1"
	"os"
)

func (s *OssutilCommandSuite) TestDemoTest(c *C) {
	c.Log("测试 commit")
}


func (s *OssutilCommandSuite) TestDemoTestDemo(c *C) {
	configFile1 := "ossutil-config-"+randLowStr(4)
	level := "info"
	level2 := "debug"
	data := "[Credentials]" + "\n" + "language=" + DefaultLanguage + "\n" + "accessKeyID=" + accessKeyID + "\n" + "accessKeySecret=" + accessKeySecret + "\n" + "endpoint=" +
		endpoint + "\n" + "loglevel=" + level2 + "\n" + "[Default]" + "\n" + "loglevel=" + level + "\n"

	s.createFile(configFile1, data, c)

	os.Args = []string{"ossutil","ls","oss://demo-walker-1889","--loglevel=info"}
	commandLine = getCommandLine()

	clearEnv()

	args, options, err := ParseArgOptions()
	c.Assert(err,IsNil)
	testLogger.Print(args)
	testLogger.Print(options)
	/*err := ParseAndRunCommand()
	c.Assert(err,IsNil)


	f, err := os.Stat(logName)
	c.Assert(err, IsNil)
	c.Assert(f.Size() > 0, Equals, true)
	strContent := s.readFile(logName, c)
	testLogger.Print(strContent)

	time.Sleep(5*time.Second)
	os.Args = []string{"ossutil","ls","oss://demo-walker-6961","--loglevel=debug","--config-file="+configFile1}
	err = ParseAndRunCommand()
	c.Assert(err,IsNil)

	f, err = os.Stat(logName)
	c.Assert(err, IsNil)
	c.Assert(f.Size() > 0, Equals, true)
	strContent = s.readFile(logName, c)
	testLogger.Print(strContent)*/

}


