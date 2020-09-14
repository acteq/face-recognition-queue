package main

import (
	"testing"
	"strings"
	"strconv"
	"bytes"
)

func TestInitEngine(t *testing.T) {

	handle := initEngine()
	
	if handle == nil {
		t.Errorf("initEngine features failed")
	}
}

func TestExtract(t *testing.T) {
	filePath := "face-test.jpg"
	featureStr := "0 128 250 68 0 0 160 65 192 75 193 188 247 97 131 189 110 2 2 62 205 127 186 61 232 47 120 189 171 83 77 60 165 60 210 189 183 66 187 60 102 7 9 61 103 87 30 189 181 246 109 189 85 90 100 61 145 166 223 189 117 219 193 188 26 253 48 190 196 232 28 58 48 31 15 61 226 248 188 59 113 64 199 61 225 154 209 61 65 13 201 188 86 10 97 189 86 127 120 61 229 201 183 189 68 210 76 189 162 125 157 189 200 180 207 187 178 196 195 189 99 83 205 188 154 212 104 61 122 191 16 189 26 48 162 61 145 105 241 189 173 146 240 187 219 166 142 61 191 173 244 61 100 174 225 189 252 207 153 61 231 74 150 189 127 234 100 188 73 50 60 61 53 202 0 61 23 160 249 59 206 144 57 188 102 213 132 61 31 249 192 188 112 140 131 188 249 224 152 189 19 194 134 188 214 176 121 61 37 89 50 59 38 178 122 189 126 116 63 187 47 195 96 59 70 206 30 61 183 119 159 61 104 77 12 61 0 195 19 61 134 172 133 188 220 224 183 61 28 129 45 189 112 21 54 189 11 75 224 189 44 60 159 187 133 44 229 59 137 110 52 61 81 110 46 188 219 153 238 61 92 205 148 189 164 84 153 187 61 61 8 61 117 230 247 59 110 64 22 61 84 149 248 189 163 234 19 190 40 6 222 61 63 71 245 60 239 103 177 61 33 35 194 186 196 49 41 60 236 5 178 61 37 97 117 61 67 208 156 61 240 233 184 189 93 237 210 60 213 92 187 189 92 155 21 61 80 245 169 189 236 41 196 60 56 17 231 188 179 239 15 60 176 104 134 60 104 121 18 189 15 54 144 189 139 71 251 60 225 196 75 190 104 192 14 188 217 29 135 188 218 223 236 188 43 70 70 189 197 26 137 189 51 152 56 61 149 217 234 60 175 86 58 189 4 105 228 188 219 148 84 61 61 146 55 60 85 13 131 61 208 63 181 189 107 111 116 61 211 110 115 189 206 0 27 188 112 113 137 59 234 61 205 59 37 254 8 62 161 55 188 60 230 31 86 62 33 192 26 61 245 169 253 188 129 232 182 61 205 158 12 190 58 203 132 189 196 159 174 61 181 155 151 61 56 184 188 61 240 93 246 189 10 96 31 189 206 15 183 188 219 139 35 61 112 67 138 61 24 166 232 61 231 128 5 60 177 243 44 189 159 47 160 61 53 241 189 188 4 96 235 60 94 80 220 60 31 200 36 188 157 225 197 61 147 51 174 188 188 107 154 189 183 129 140 189 63 172 21 189 252 160 18 59 14 184 125 60 171 51 245 188 86 106 244 189 7 51 160 60 6 47 74 189 50 137 244 60 203 106 161 60 205 121 62 189 72 232 159 61 7 113 171 61 10 76 188 188 88 228 229 189 114 73 117 60 63 128 150 189 4 0 146 189 141 93 21 187 9 204 202 188 27 124 75 60 229 132 178 61 247 46 159 189 171 146 215 187 131 61 76 189 171 209 75 61 69 48 172 186 8 77 237 59 37 141 27 189 200 153 36 189 194 188 163 61 97 129 195 188 81 129 239 189 70 81 189 188 35 113 55 60 44 210 175 187 41 63 157 189 232 0 33 190 14 156 242 61 40 64 136 188 199 118 56 189 245 159 183 188 139 74 151 189 169 29 10 61 148 197 128 61 41 220 35 189 27 48 247 60 220 146 130 187 181 141 193 60 191 234 184 188 85 226 200 61 80 107 255 61 93 50 90 189 64 69 86 61 6 170 161 58 91 12 118 61 169 4 12 61 106 222 113 61 125 211 190 60 21 200 56 189 247 2 205 189 245 194 194 189 113 243 53 61 18 150 66 189 57 127 88 61 79 82 130 189 5 23 252 57 27 224 4 62 165 241 4 61 36 196 179 61 156 140 70 60 168 137 94 61 128 110 65 61 92 118 33 62 14 165 129 60 110 169 148 61 98 139 89 61 234 44 132 189 3 0 0 189 135 87 170 59 95 246 137 188 190 70 236 189 253 68 145 189 115 91 145 189 174 225 203 188 44 13 63 60 80 34 201 188 186 39 187 60 48 84 231 58 117 122 44 61 251 188 171 188 80 4 234 187 250 74 215 61 244 85 116 189 33 116 89 189 74 80 12 61 1 18 198 61 80 92 88 189 88 189 187 189 58 24 139 189 181 71 202 61 199 203 137 189 104 242 40 61 39 86 31 61 36 166 8 62 171 82 167 189 66 27 157 61 164 235 130 189 223 216 118 59 170 39 124 59 60 205 196 189 253 174 228 60 197 174 178 187 241 36 218 60 242 111 218 189"
	
	strArray := strings.Split(featureStr, " ")
	featureBytes := make([]byte, len(strArray))
	for i, c := range(strArray) {
	    val, _ := strconv.ParseInt(c, 10, 32)
		featureBytes[i] = byte(val)
	}
	handle := initEngine()
	results := extract(filePath, handle)

	if ! bytes.Equal(featureBytes, results) {
		t.Errorf("extract features failed: feature length = %d, result length = %d", len(results), len(results) )
	}
}

