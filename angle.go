package si

import "math"

const (
	Radian         Angle = 1e0
	Yoctoradian          = Yocto * Radian
	Zeptoradian          = Zepto * Radian
	Attoradian           = Atto * Radian
	Femtoradian          = Femto * Radian
	Picoradian           = Pico * Radian
	Nanoradian           = Nano * Radian
	Microradian          = Micro * Radian
	Milliradian          = Milli * Radian
	Centiradian          = Centi * Radian
	Deciradian           = Deci * Radian
	Degree               = Radian * math.Pi / 180
	Arcminute            = Degree / 60
	Arcsecond            = Degree / 3600
	Milliarcsecond       = Milli * Arcsecond
	Microarcsecond       = Micro * Arcsecond
)
