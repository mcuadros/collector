package collector

import (
	. "collector/logger"
	"runtime"
	"time"
)

type Collector struct {
	writer  *Writer
	reader  *Reader
	channel chan map[string]string
}

func NewCollector() *Collector {
	collector := new(Collector)

	return collector
}

func (self *Collector) Configure(filename string) {
	GetConfig().LoadFile(filename)
}

func (self *Collector) Boot() {
	self.configureLogger()
	self.configureMaxProcs()
	self.bootWriter()
	self.bootReader()
}

func (self *Collector) configureLogger() {
	Info("Starting ...")
}

func (self *Collector) configureMaxProcs() {
	Info("Number of max. process %d", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func (self *Collector) bootWriter() {
	self.writer = GetContainer().GetWriter()
}

func (self *Collector) bootReader() {
	self.reader = GetContainer().GetReader()
}

func (self *Collector) Run() {
	self.channel = self.writer.GoWriteFromChannel()
	self.reader.GoReadIntoChannel(self.channel)
	self.wait()
	self.reader.Finish()
}

func (self *Collector) wait() {
	for self.writer.IsAlive() {
		time.Sleep(1 * time.Second)
		self.writer.PrintCounters(1)
	}

	Info("nothing more for read, terminating daemon ...")
}
