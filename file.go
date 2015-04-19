package utility

import "os"
import "io/ioutil"
import log "code.google.com/p/log4go"
import proto "code.google.com/p/goprotobuf/proto"

func LoadFile(file string, pb proto.Message) {
	f, err := os.Open(file)
	if err != nil {
		log.Error("Load file %s error %v\n", file, err)
		os.Exit(1)
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error("Read content error %v\n", err)
		os.Exit(1)
	}

	err = proto.UnmarshalText(string(content), pb)
	if err != nil {
		log.Error("UnmarshalText content error %v\n", err)
		os.Exit(1)
	}

	log.Fine("load %s success info:%s\n", file, pb.String())
}
