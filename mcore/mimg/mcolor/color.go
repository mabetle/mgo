package mcolor

import (
	ic "image/color"
)

type RGB struct {
	r, g, b int
}

func NewRGB(r, g, b int) RGB {
	return RGB{r, g, b}
}

func (rgb RGB) RGBA(a int) ic.Color {
	return ic.RGBA{uint8(rgb.r), uint8(rgb.g), uint8(rgb.b), uint8(a)}
}

// RGB with
func RGBA(r, g, b int) ic.Color {
	return ic.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}
}

func GetRGB(c ic.Color) RGB {
	r, g, b, _ := c.RGBA()
	return NewRGB(int(r), int(g), int(b))
}

var (
	Black = RGBA(0, 0, 0)
	// black（黑）
	Navy = RGBA(0, 0, 128)
	// navy（海军蓝）
	DaryBlue = RGBA(0, 0, 139)
	// darkblue（深蓝）
	MediumBlue = RGBA(0, 0, 205)
	// mediumblue（适中的蓝）
	Blue = RGBA(0, 0, 255)
	// blue（蓝）
	Teal = RGBA(0, 128, 128)
	// teal（水鸭色）
	DarkCyan = RGBA(0, 139, 139)
	// darkcyan（深青色）
	DeepskyBlue = RGBA(0, 191, 255)
	// deepskyblue（深天蓝）
	DarkturQuoise = RGBA(0, 206, 209)
	// darkturquoise（深宝石绿）
	DarkslateGray = RGBA(47, 79, 79)
	// darkslategray（深石板灰）
	Lime = RGBA(0, 255, 0)
	// lime（柠檬）
	SprintGreen = RGBA(0, 255, 127)
	// springgreen（春天绿）
	Aqua = RGBA(0, 255, 255)
	// aqua（水绿）
	CornFlowerBlue = RGBA(100, 149, 237)
	// cornflowerblue（矢车菊蓝）
	LightSteelBlue = RGBA(176, 196, 222)
	// lightsteelblue（淡钢蓝）
	LightSlateGray = RGBA(119, 136, 153)
	// lightslategray（浅石板灰）
	SlateGray = RGBA(112, 128, 144)
	// slategray（石板灰）
	MediumAquaMarine = RGBA(102, 205, 170)
	// mediumaquamarine（适中的碧绿）
	MidiumSpringGreen = RGBA(0, 250, 154)
	// mediumspringgreen（适中的春天绿）
	MintCream = RGBA(245, 255, 250)
	// mintcream（薄荷奶油）
	DimGray = RGBA(105, 105, 105)
	// dimgray（暗灰）
	OliverDrab = RGBA(107, 142, 35)
	// olivedrab（橄榄褐）
	MediumSlateBlue = RGBA(123, 104, 238)
	// mediumslateblue（适中的的板岩蓝）
	SlateBlue = RGBA(106, 90, 205)
	// slateblue（板岩蓝）
	LawnGreen = RGBA(124, 252, 0)
	// lawngreen（草坪绿）
	ChartReuse = RGBA(127, 255, 0)
	// chartreuse（查特酒绿）
	AquaMarine = RGBA(127, 255, 212)
	// aquamarine（碧绿）
	Maroon = RGBA(128, 0, 0)
	// maroon（粟色）
	Purple = RGBA(128, 0, 128)
	// purple（紫）
	Olive = RGBA(128, 128, 0)
	// olive（橄榄）
	SkyBlue = RGBA(135, 206, 235)
	// skyblue（天蓝）
	LightSkyBlue = RGBA(135, 206, 250)
	// lightskyblue（淡天蓝）
	BlueViolet = RGBA(138, 43, 226)
	// blueviolet（蓝紫罗兰）
	DarkRed = RGBA(139, 0, 0)
	// darkred（深红）
	DarkMagenta = RGBA(139, 0, 139)
	// darkmagenta（深洋紫）
	DarkSeaGreen = RGBA(143, 188, 143)
	// darkseagreen（深海洋绿）
	LimeGreen = RGBA(50, 205, 50)
	// limegreen（柠檬绿）
	LightGreen = RGBA(144, 238, 144)
	// lightgreen（浅绿）
	MidiumPurple = RGBA(147, 112, 219)
	// mediumpurple（适中的紫）
	PaleGreen = RGBA(152, 251, 152)
	// palegreen（苍白绿）
	Sienna = RGBA(160, 82, 45)
	// sienna（土黄赭）
	Brown = RGBA(165, 42, 42)
	// brown（棕）
	DaryGray = RGBA(169, 169, 169)
	// darkgray（深灰）
	LightBlue = RGBA(173, 216, 230)
	// lightblue（淡蓝）
	GreenYellow = RGBA(173, 255, 47)
	// greenyellow（绿黄）
	PaleturQuoise = RGBA(175, 238, 238)
	// paleturquoise（苍白的宝石绿）
	Cyan = RGBA(0, 255, 255)
	// cyan（青）
	PowderBlue = RGBA(176, 224, 230)
	// powderblue（火药蓝）
	Firebrick = RGBA(178, 34, 34)
	// firebrick（火砖）
	DarkGoldenRod = RGBA(184, 134, 11)
	// darkgoldenrod（深秋）
	MediumOrchid = RGBA(186, 85, 211)
	// mediumorchid（适中的兰花紫）
	DarkViolet = RGBA(148, 0, 211)
	// darkviolet（深紫罗兰）
	RosyBrown = RGBA(188, 143, 143)
	// rosybrown（玫瑰棕）
	DarkKhaki = RGBA(189, 183, 107)
	// darkkhaki（深卡其布）
	ChocolateSaddleBrown = RGBA(192, 14, 235)
	// chocolatesaddlebrown（马鞍棕）
	SeaShell = RGBA(255, 245, 238)
	// seashell（海贝）
	Sliver = RGBA(192, 192, 192)
	// silver（银白）
	MediumVioletRed = RGBA(199, 21, 133)
	// mediumvioletred（适中的紫罗兰红）
	Orchid = RGBA(218, 112, 214)
	// orchid（兰花紫）
	Peru = RGBA(205, 133, 63)
	// peru（秘鲁）
	IndianRed = RGBA(205, 92, 92)
	// indianred（浅粉红）
	Chocolate = RGBA(210, 105, 30)
	// chocolate（巧克力）
	Tan = RGBA(210, 180, 140)
	// tan（晒）
	Thistle = RGBA(216, 191, 216)
	// thistle（苍紫）
	GoldenRod = RGBA(218, 165, 32)
	// goldenrod（秋）
	PaleVioletRed = RGBA(219, 112, 147)
	// palevioletred（脸红的淡紫红）
	HotPink = RGBA(255, 105, 180)
	// hotpink（热情的粉红）
	Crimson = RGBA(220, 20, 60)
	// crimson（腥红）
	GainBoro = RGBA(220, 220, 220)
	// gainsboro（赶死部落）
	LightGray = RGBA(211, 211, 211)
	// lightgrey（浅灰）
	Plum = RGBA(221, 160, 221)
	// plum（轻紫）
	BurlyWood = RGBA(222, 184, 135)
	// burlywood（树干）
	LightCyan = RGBA(224, 255, 255)
	// lightcyan（淡青）
	Lavender = RGBA(230, 230, 250)
	// lavender（熏衣草花的淡紫）
	Violet = RGBA(238, 130, 238)
	// violet（紫罗兰）
	LightCoral = RGBA(240, 128, 128)
	// lightcoral（浅珊瑚）
	Khaki = RGBA(240, 230, 140)
	// khaki（卡其布）
	AliceBlue = RGBA(240, 248, 255)
	// aliceblue（爱丽丝蓝）
	HoneyDew = RGBA(240, 255, 240)
	// honeydew（浅粉红）
	Azure = RGBA(240, 255, 255)
	// azure（蔚蓝）
	SandyBrown = RGBA(244, 164, 96)
	// sandybrown（沙棕）
	Wheat = RGBA(245, 222, 179)
	// wheat（小麦）
	Beige = RGBA(245, 245, 220)
	// beige（米色）
	WhiteSmoke = RGBA(245, 245, 245)
	// whitesmoke（烟白）
	MidNightBlue = RGBA(25, 25, 112)
	// midnightblue（午夜蓝）
	Linen = RGBA(250, 240, 230)
	// linen（亚麻）
	LightGoldenRodYello = RGBA(250, 250, 210)
	// lightgoldenrodyellow（浅秋黄）
	Ivory = RGBA(255, 255, 240)
	// ivory（象牙白）
	Oldlace = RGBA(253, 245, 230)
	// oldlace（浅米色）
	Red = RGBA(255, 0, 0)
	// red（红）
	Fuchsia = RGBA(255, 0, 255)
	// fuchsia（紫红）
	Magenta = RGBA(255, 0, 255)
	// magenta（洋紫)
	Coral = RGBA(255, 127, 80)
	// coral（珊瑚）
	DarkOrange = RGBA(255, 140, 0)
	// darkorange（深橙色）
	LightSalmon = RGBA(255, 160, 122)
	// lightsalmon（浅肉）
	Orange = RGBA(255, 165, 0)
	// orange（橙）
	Pink = RGBA(255, 192, 203)
	// pink（粉红）
	Gold = RGBA(255, 215, 0)
	// gold（金）
	Moccasin = RGBA(255, 228, 181)
	// moccasin（鹿皮）
	Bisque = RGBA(255, 228, 196)
	// bisque（乳脂）
	MistyRose = RGBA(255, 228, 225)
	// mistyrose（雾中玫瑰）
	Salmon = RGBA(250, 128, 114)
	// salmon（肉）
	BlancheDalmond = RGBA(255, 235, 205)
	// blanchedalmond（漂白后的杏仁）
	NavajoWhite = RGBA(255, 222, 173)
	// navajowhite（耐而节白）
	AntiqueWhite = RGBA(250, 235, 215)
	// antiquewhite（古白）
	PapaYawhip = RGBA(255, 239, 213)
	// papayawhip（木瓜）

	LavenderBlush = RGBA(255, 240, 245)
	// lavenderblush（苍白的紫罗兰红）
	CronSilk = RGBA(255, 248, 220)
	// cornsilk（玉米）
	LemonChiffon = RGBA(255, 250, 205)
	// lemonchiffon（柠檬沙）
	PaleGoldenRod = RGBA(238, 232, 170)
	// palegoldenrod（灰秋）
	FloralWhite = RGBA(255, 250, 240)
	// floralwhite（白花）
	Yellow = RGBA(255, 255, 0)
	// yellow（黄）
	LightYellow = RGBA(255, 255, 224)
	// lightyellow（浅黄）
	White      = RGBA(255, 255, 255)
	OrgangeRed = RGBA(255, 69, 0)
	// orangered（橙红）
	Tomato = RGBA(255, 99, 71)
	// tomato（番茄色）
	DodgeBlue = RGBA(30, 144, 255)
	// dodgerblue（道奇蓝）
	ForestGreen = RGBA(34, 139, 34)
	// forestgreen（森林绿）
	MediumGreen = RGBA(60, 179, 113)
	// mediumseagreen（适中的海洋绿）
	SeaGreen = RGBA(46, 139, 87)
	// seagreen（海洋绿）
	RoyalBlue = RGBA(65, 105, 225)
	// royalblue（皇家蓝）
	SteelBlue = RGBA(70, 130, 180)
	// steelblue（钢蓝）
	MidiumTurquoise = RGBA(72, 209, 204)
	// mediumturquoise（适中的宝石绿）
	LightseaGreen = RGBA(32, 178, 170)
	// lightseagreen（浅海洋绿）
	Turquoise = RGBA(64, 224, 208)
	// turquoise（宝石绿）
	DarkslateBlue = RGBA(72, 61, 139)
	// darkslateblue（深板岩蓝）
	Indigo = RGBA(75, 0, 130)
	// indigo（靓青）
	DarkkoliveGreen = RGBA(85, 107, 47)
	// darkolivegreen（深橄榄绿）
	YellowGreen = RGBA(154, 205, 50)
	// yellowgreen（黄绿）
	CadeBlue = RGBA(95, 158, 160)
	// cadetblue（军校蓝）
)
