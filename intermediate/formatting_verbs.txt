package intermediate

import "fmt"

func main() {

	// --- General Formatting Verbs
	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign

	// i := 111_505.5
	// string := "Hello World!"

	// fmt.Printf("%v\n", i)
	// fmt.Printf("%#v\n", i)
	// fmt.Printf("%T\n", i)
	// fmt.Printf("%v%%\n", i)

	// fmt.Printf("%v\n", string)
	// fmt.Printf("%#v\n", string)
	// fmt.Printf("%T\n", string)

	// --- Integer Formatting Verbs
	// 	%b	Base 2
	// %d	Base 10
	// %+d	Base 10 and always show sign
	// %o	Base 8
	// %O	Base 8, with leading 0o
	// %x	Base 16, lowercase
	// %X	Base 16, uppercase
	// %#x	Base 16, with leading 0x
	// %4d	Pad with spaces (width 4, right justified)
	// %-4d	Pad with spaces (width 4, left justified)
	// %04d	Pad with zeroes (width 4

	// int := 255

	// fmt.Printf("%b\n", int)
	// fmt.Printf("%d\n", int)
	// fmt.Printf("%+d\n", int)
	// fmt.Printf("%o\n", int)
	// fmt.Printf("%O\n", int)
	// fmt.Printf("%x\n", int)
	// fmt.Printf("%X\n", int)
	// fmt.Printf("%#x\n", int)
	// fmt.Printf("%4d\n", int)
	// fmt.Printf("%-4d\n", int)
	// fmt.Printf("%04d\n", int)

	// --- String Formatting Verbs
	// %s	Prints the value as plain string
	// %q	Prints the value as a double-quoted string
	// %8s	Prints the value as plain string (width 8, right justified)
	// %-8s	Prints the value as plain string (width 8, left justified)
	// %x	Prints the value as hex dump of byte values
	// % x	Prints the value as hex dump with spaces

	// txt := "World"

	// fmt.Printf("%s\n", txt)
	// fmt.Printf("%q\n", txt)
	// fmt.Printf("%8s\n", txt)
	// fmt.Printf("%-8s\n", txt)
	// fmt.Printf("%x\n", txt)
	// fmt.Printf("% x\n", txt)

	// --- Boolean Formatting Verbs
	// %t	Value of the boolean operator in true or false format (same as using %v)

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", t)
	fmt.Printf("%t\n", f)

	// 	--- Float Formatting Verbs
	// Verb	Description
	// %e	Scientific notation with 'e' as exponent
	// %f	Decimal point, no exponent
	// %.2f	Default width, precision 2
	// %6.2f Minimum Width 6, precision 2
	// %g	Exponent as needed, only necessary digits

	flt := 9180000.00

	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)

}