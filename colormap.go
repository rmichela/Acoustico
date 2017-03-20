package main

import "math"

// Colormap is a function that returns a Red, Green, and Blue value for an input
// The input must be between [0..1)
type Colormap func(float64) (byte, byte, byte)

// Inferno converts an input value to the Inferno colormap from python's matplotlib
func Inferno(value float64) (r byte, g byte, b byte) {
	// clamp output
	if value < 0 {
		return 0, 0, 0
	}
	if value >= 1.0 {
		return 255, 255, 255
	}
	// scale value to [0..256)
	i := math.Floor(value * 256.0)
	switch i {
	case 0:
		return 0, 0, 3
	case 1:
		return 0, 0, 4
	case 2:
		return 0, 0, 6
	case 3:
		return 1, 0, 7
	case 4:
		return 1, 1, 9
	case 5:
		return 1, 1, 11
	case 6:
		return 2, 1, 14
	case 7:
		return 2, 2, 16
	case 8:
		return 3, 2, 18
	case 9:
		return 4, 3, 20
	case 10:
		return 4, 3, 22
	case 11:
		return 5, 4, 24
	case 12:
		return 6, 4, 27
	case 13:
		return 7, 5, 29
	case 14:
		return 8, 6, 31
	case 15:
		return 9, 6, 33
	case 16:
		return 10, 7, 36
	case 17:
		return 12, 7, 38
	case 18:
		return 13, 8, 40
	case 19:
		return 14, 8, 43
	case 20:
		return 15, 9, 45
	case 21:
		return 16, 9, 47
	case 22:
		return 18, 10, 50
	case 23:
		return 19, 10, 52
	case 24:
		return 20, 11, 55
	case 25:
		return 22, 11, 57
	case 26:
		return 23, 11, 59
	case 27:
		return 25, 11, 62
	case 28:
		return 26, 12, 64
	case 29:
		return 28, 12, 67
	case 30:
		return 29, 12, 69
	case 31:
		return 31, 12, 72
	case 32:
		return 33, 12, 74
	case 33:
		return 34, 11, 76
	case 34:
		return 36, 11, 78
	case 35:
		return 38, 11, 81
	case 36:
		return 39, 11, 83
	case 37:
		return 41, 11, 85
	case 38:
		return 43, 10, 87
	case 39:
		return 45, 10, 89
	case 40:
		return 46, 10, 90
	case 41:
		return 48, 10, 92
	case 42:
		return 50, 9, 94
	case 43:
		return 52, 9, 95
	case 44:
		return 54, 9, 96
	case 45:
		return 55, 9, 98
	case 46:
		return 57, 9, 99
	case 47:
		return 59, 9, 100
	case 48:
		return 60, 9, 101
	case 49:
		return 62, 9, 102
	case 50:
		return 64, 9, 103
	case 51:
		return 66, 9, 104
	case 52:
		return 67, 10, 104
	case 53:
		return 69, 10, 105
	case 54:
		return 71, 10, 106
	case 55:
		return 72, 11, 106
	case 56:
		return 74, 11, 107
	case 57:
		return 76, 12, 107
	case 58:
		return 77, 12, 108
	case 59:
		return 79, 13, 108
	case 60:
		return 80, 13, 108
	case 61:
		return 82, 14, 109
	case 62:
		return 84, 14, 109
	case 63:
		return 85, 15, 109
	case 64:
		return 87, 15, 109
	case 65:
		return 89, 16, 110
	case 66:
		return 90, 17, 110
	case 67:
		return 92, 17, 110
	case 68:
		return 93, 18, 110
	case 69:
		return 95, 18, 110
	case 70:
		return 97, 19, 110
	case 71:
		return 98, 20, 110
	case 72:
		return 100, 20, 110
	case 73:
		return 101, 21, 110
	case 74:
		return 103, 21, 110
	case 75:
		return 104, 22, 110
	case 76:
		return 106, 23, 110
	case 77:
		return 108, 23, 110
	case 78:
		return 109, 24, 110
	case 79:
		return 111, 24, 110
	case 80:
		return 112, 25, 110
	case 81:
		return 114, 26, 110
	case 82:
		return 116, 26, 110
	case 83:
		return 117, 27, 110
	case 84:
		return 119, 27, 109
	case 85:
		return 120, 28, 109
	case 86:
		return 122, 28, 109
	case 87:
		return 124, 29, 109
	case 88:
		return 125, 29, 108
	case 89:
		return 127, 30, 108
	case 90:
		return 128, 31, 108
	case 91:
		return 130, 31, 108
	case 92:
		return 132, 32, 107
	case 93:
		return 133, 32, 107
	case 94:
		return 135, 33, 107
	case 95:
		return 136, 33, 106
	case 96:
		return 138, 34, 106
	case 97:
		return 140, 35, 105
	case 98:
		return 141, 35, 105
	case 99:
		return 143, 36, 104
	case 100:
		return 144, 36, 104
	case 101:
		return 146, 37, 104
	case 102:
		return 148, 37, 103
	case 103:
		return 149, 38, 103
	case 104:
		return 151, 39, 102
	case 105:
		return 152, 39, 101
	case 106:
		return 154, 40, 101
	case 107:
		return 155, 40, 100
	case 108:
		return 157, 41, 100
	case 109:
		return 159, 42, 99
	case 110:
		return 160, 42, 98
	case 111:
		return 162, 43, 98
	case 112:
		return 163, 43, 97
	case 113:
		return 165, 44, 96
	case 114:
		return 167, 45, 96
	case 115:
		return 168, 45, 95
	case 116:
		return 170, 46, 94
	case 117:
		return 171, 47, 93
	case 118:
		return 173, 47, 93
	case 119:
		return 174, 48, 92
	case 120:
		return 176, 49, 91
	case 121:
		return 177, 49, 90
	case 122:
		return 179, 50, 89
	case 123:
		return 180, 51, 89
	case 124:
		return 182, 52, 88
	case 125:
		return 183, 52, 87
	case 126:
		return 185, 53, 86
	case 127:
		return 186, 54, 85
	case 128:
		return 188, 55, 84
	case 129:
		return 189, 56, 83
	case 130:
		return 191, 56, 82
	case 131:
		return 192, 57, 81
	case 132:
		return 194, 58, 80
	case 133:
		return 195, 59, 79
	case 134:
		return 197, 60, 78
	case 135:
		return 198, 61, 77
	case 136:
		return 199, 62, 76
	case 137:
		return 201, 63, 75
	case 138:
		return 202, 64, 74
	case 139:
		return 203, 65, 73
	case 140:
		return 205, 66, 72
	case 141:
		return 206, 67, 71
	case 142:
		return 207, 68, 70
	case 143:
		return 209, 69, 69
	case 144:
		return 210, 70, 68
	case 145:
		return 211, 71, 67
	case 146:
		return 213, 72, 65
	case 147:
		return 214, 73, 64
	case 148:
		return 215, 74, 63
	case 149:
		return 216, 76, 62
	case 150:
		return 217, 77, 61
	case 151:
		return 219, 78, 60
	case 152:
		return 220, 79, 59
	case 153:
		return 221, 81, 57
	case 154:
		return 222, 82, 56
	case 155:
		return 223, 83, 55
	case 156:
		return 224, 85, 54
	case 157:
		return 225, 86, 53
	case 158:
		return 226, 87, 51
	case 159:
		return 227, 89, 50
	case 160:
		return 228, 90, 49
	case 161:
		return 229, 91, 48
	case 162:
		return 230, 93, 47
	case 163:
		return 231, 94, 45
	case 164:
		return 232, 96, 44
	case 165:
		return 233, 97, 43
	case 166:
		return 234, 99, 42
	case 167:
		return 235, 100, 40
	case 168:
		return 236, 102, 39
	case 169:
		return 237, 103, 38
	case 170:
		return 237, 105, 37
	case 171:
		return 238, 106, 35
	case 172:
		return 239, 108, 34
	case 173:
		return 240, 110, 33
	case 174:
		return 241, 111, 32
	case 175:
		return 241, 113, 30
	case 176:
		return 242, 114, 29
	case 177:
		return 243, 116, 28
	case 178:
		return 243, 118, 26
	case 179:
		return 244, 119, 25
	case 180:
		return 244, 121, 24
	case 181:
		return 245, 123, 22
	case 182:
		return 246, 125, 21
	case 183:
		return 246, 126, 20
	case 184:
		return 247, 128, 18
	case 185:
		return 247, 130, 17
	case 186:
		return 248, 132, 16
	case 187:
		return 248, 133, 14
	case 188:
		return 248, 135, 13
	case 189:
		return 249, 137, 12
	case 190:
		return 249, 139, 11
	case 191:
		return 250, 141, 9
	case 192:
		return 250, 142, 8
	case 193:
		return 250, 144, 8
	case 194:
		return 251, 146, 7
	case 195:
		return 251, 148, 6
	case 196:
		return 251, 150, 6
	case 197:
		return 251, 152, 6
	case 198:
		return 252, 153, 6
	case 199:
		return 252, 155, 6
	case 200:
		return 252, 157, 6
	case 201:
		return 252, 159, 7
	case 202:
		return 252, 161, 7
	case 203:
		return 252, 163, 8
	case 204:
		return 252, 165, 10
	case 205:
		return 252, 167, 11
	case 206:
		return 252, 169, 13
	case 207:
		return 252, 170, 14
	case 208:
		return 252, 172, 16
	case 209:
		return 252, 174, 18
	case 210:
		return 252, 176, 20
	case 211:
		return 252, 178, 22
	case 212:
		return 252, 180, 24
	case 213:
		return 252, 182, 26
	case 214:
		return 252, 184, 28
	case 215:
		return 252, 186, 30
	case 216:
		return 251, 188, 33
	case 217:
		return 251, 190, 35
	case 218:
		return 251, 192, 37
	case 219:
		return 251, 194, 40
	case 220:
		return 250, 196, 42
	case 221:
		return 250, 198, 45
	case 222:
		return 250, 200, 47
	case 223:
		return 249, 202, 50
	case 224:
		return 249, 204, 52
	case 225:
		return 249, 206, 55
	case 226:
		return 248, 208, 58
	case 227:
		return 248, 210, 61
	case 228:
		return 247, 212, 63
	case 229:
		return 247, 214, 66
	case 230:
		return 246, 216, 69
	case 231:
		return 246, 217, 73
	case 232:
		return 245, 219, 76
	case 233:
		return 245, 221, 79
	case 234:
		return 244, 223, 82
	case 235:
		return 244, 225, 86
	case 236:
		return 244, 227, 89
	case 237:
		return 243, 229, 93
	case 238:
		return 243, 231, 97
	case 239:
		return 242, 233, 101
	case 240:
		return 242, 234, 105
	case 241:
		return 242, 236, 109
	case 242:
		return 242, 238, 113
	case 243:
		return 242, 239, 117
	case 244:
		return 242, 241, 121
	case 245:
		return 242, 243, 125
	case 246:
		return 243, 244, 130
	case 247:
		return 243, 245, 134
	case 248:
		return 244, 247, 138
	case 249:
		return 245, 248, 142
	case 250:
		return 246, 249, 146
	case 251:
		return 247, 251, 150
	case 252:
		return 248, 252, 154
	case 253:
		return 249, 253, 157
	case 254:
		return 251, 254, 161
	}
	return 0, 0, 0
}
