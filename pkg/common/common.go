package common

const (
	TIME_FORMAT = "02 Jan 2006"
)

const (
	SUCCESS        = 0
	FAILED         = 1
	NOT_OPEN       = 2
	NOT_READ       = 3
	NOT_WRITE      = 4
	NOT_MARSHALL   = 5
	NOT_UNMARSHALL = 6
)

func MKFAIL(erCode int) int {
	if erCode > 0 {
		return -erCode
	}
	return erCode
}

func MKSUCCESS() int {
	return SUCCESS
}

func IS_SUCCESS(erCode int) bool {
	return erCode >= 0
}
