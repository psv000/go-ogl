module sample

require (
	framework v0.1.0
	github.com/sirupsen/logrus v1.4.2
)

replace framework => ../framework

go 1.13
