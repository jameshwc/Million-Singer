package mysql

func init() {

}

func escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var esc byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		esc = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			esc = '0'
			break
		case '\n': /* Must be escaped for logs */
			esc = 'n'
			break
		case '\r':
			esc = 'r'
			break
		case '\\':
			esc = '\\'
			break
		case '\'':
			esc = '\''
			break
		case '"': /* Better safe than sorry */
			esc = '"'
			break
		case '\032': /* This gives problems on Win32 */
			esc = 'Z'
		}

		if esc != 0 {
			dest = append(dest, '\\', esc)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}
