package str

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestAfter(t *testing.T) {

	check := func(subject, search, expected string) {
		actual := After(subject, search)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("hannah", "han", "nah")
	check("hannah", "n", "nah")
	check("ééé hannah", "han", "nah")
	check("ééé hannah xxx", "han", "nah xxx")
	check("hannah", "xxxx", "hannah")
	check("hannah", "", "hannah")
	check("han0nah", "0", "nah")
	check("chris", "chr", "is")
	check("chris", "c", "hris")
	check("ééé chris", "r", "is")
	check("ééé chris ééé chris", "s", " ééé chris")
	check("", "a", "")
	check("@$@£^$%&%$", "^", "$%&%$")
	check("@$ @£^$% & %  $", "£", "^$% & %  $")
	check("12345", "4", "5")
}

func TestAfterLast(t *testing.T) {

	check := func(subject, search, expected string) {
		actual := AfterLast(subject, search)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("yvette", "yve", "tte")
	check("yvette", "t", "e")
	check("ééé yvette", "t", "e")
	check("yvette", "tte", "")
	check("yvette", "xxxx", "yvette")
	check("yvette", "", "yvette")
	check("yv0et0te", "0", "te")
	check("yv2et2te", "2", "te")
	check("----foo", "---", "foo")
	check("聚合科技有限公司", "科技", "有限公司")
	check("聚合公司--有限公司科技", "公司", "科技")
}

func TestBefore(t *testing.T) {

	check := func(subject, search, expected string) {
		actual := Before(subject, search)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("hannah", "nah", "han")
	check("hannah", "n", "ha")
	check("ééé hannah", "han", "ééé ")
	check("hannah", "xxxx", "hannah")
	check("hannah", "", "hannah")
	check("han0nah", "0", "han")
	check("", "", "")
	check("", "a", "")
	check("a", "a", "")
	check("foo@bar.com", "@", "foo")
	check("foo@@bar.com", "@", "foo")
	check("@foo@bar.com", "@", "")
	check("聚合科技有限公司", "科技", "聚合")
	check("聚合公司--有限公司科技", "公司", "聚合")
}

func TestBeforeLast(t *testing.T) {

	check := func(subject, search, expected string) {
		actual := BeforeLast(subject, search)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("yvette", "tte", "yve")
	check("yvette", "t", "yvet")
	check("ééé yvette", "yve", "ééé ")
	check("yvette", "yve", "")
	check("yvette", "xxxx", "yvette")
	check("yvette", "", "yvette")
	check("yv0et0te", "0", "yv0et")
	check("yv2et2te", "2", "yv2et")
	check("", "test", "")
	check("yvette", "yvette", "")
	check("laravel framework", " ", "laravel")
	check("yvette\tyv0et0te", "\t", "yvette")
	check("聚合科技有限公司", "科技", "聚合")
	check("聚合公司--有限公司科技", "公司", "聚合公司--有限")
}

func TestBetween(t *testing.T) {

	check := func(subject, from, to, expected string) {
		actual := Between(subject, from, to)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("abc", "", "c", "abc")
	check("abc", "a", "", "abc")
	check("abc", "", "", "abc")
	check("abc", "a", "c", "b")
	check("dddabc", "a", "c", "b")
	check("abcddd", "a", "c", "b")
	check("dddabcddd", "a", "c", "b")
	check("hannah", "ha", "ah", "nn")
	check("[a]ab[b]", "[", "]", "a]ab[b")
	check("foofoobar", "foo", "bar", "foo")
	check("foobarbar", "foo", "bar", "bar")
	check("12345", "1", "5", "234")
	check("123456789", "123", "6789", "45")
	check("nothing", "foo", "bar", "nothing")
}

func TestBetweenFirst(t *testing.T) {

	check := func(subject, from, to, expected string) {
		actual := BetweenFirst(subject, from, to)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("abc", "", "c", "abc")
	check("abc", "a", "", "abc")
	check("abc", "", "", "abc")
	check("abc", "a", "c", "b")
	check("dddabc", "a", "c", "b")
	check("abcddd", "a", "c", "b")
	check("dddabcddd", "a", "c", "b")
	check("hannah", "ha", "ah", "nn")
	check("[a]ab[b]", "[", "]", "a")
	check("foofoobar", "foo", "bar", "foo")
	check("foobarbar", "foo", "bar", "")
}

func TestCamel(t *testing.T) {

	check := func(value, expected string) {
		actual := Camel(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("Golang_s_t_r_ing_helpers", "golangSTRIngHelpers")
	check("Golang_string_helpers", "golangStringHelpers")
	check("Golang-string-helpers", "golangStringHelpers")
	check("Golang  -_-  string   -_-   helpers   ", "golangStringHelpers")
	check("FooBar", "fooBar")
	check("foo_bar", "fooBar")
	check("Foo-barBaz", "fooBarBaz")
	check("foo-bar_baz", "fooBarBaz")
	check("", "")
	check("foo1_bar", "foo1Bar")
	check("1 foo bar", "1FooBar")
}

func TestEndsWith(t *testing.T) {

	check := func(haystack string, needles interface{}, expected bool) {
		actual := EndsWith(haystack, needles)
		if actual != expected {
			t.Errorf("Expected <%t> got <%t>", expected, actual)
		}
	}

	check("jason", "on", true)
	check("jason", "jason", true)
	check("jason", []string{"on"}, true)
	check("jason", []string{"no", "on"}, true)
	check("jason", "no", false)
	check("jason", []string{"no"}, false)
	check("jason", "", false)
	check("", "", false)
	check("jason", []string{}, false)
	check("jason", "N", false)
	check("7", " 7", false)
	check("a7", "7", true)
	// Test for multibyte string support
	check("Jönköping", "öping", true)
	check("Malmö", "mö", true)
	check("Jönköping", "oping", false)
	check("Malmö", "mo", false)
	check("你好", "好", true)
	check("你好", "你", false)
	check("你好", "a", false)
}

func TestFinish(t *testing.T) {

	check := func(value, cap, expected string) {
		actual := Finish(value, cap)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("ab", "bc", "abbc")
	check("abbcbc", "bc", "abbc")
	check("abcbbcbc", "bc", "abcbbc")
	check("1112223334456666", "6", "1112223334456")
	check("999999999", "9", "9")
	check("1112223334456666", "6666", "1112223334456666")
	check("999999999", "999", "999")
	check("12345678", "1234", "123456781234")
	check("chris", " keller", "chris keller")
	check("chris", "chris", "chris")
}

func TestIs(t *testing.T) {
	check := func(patterns interface{}, value string, expected bool) {
		actual := Is(patterns, value)
		if actual != expected {
			t.Errorf("Expected <%t> got <%t>", expected, actual)
		}
	}

	check("/", "/", true)
	check("/", " /", false)
	check("/", "/a", false)
	check("foo/*", "foo/bar/baz", true)

	check("*@*", `App\Class@method`, true)
	check("*@*", `app\Class@`, true)
	check("*@*", "@method", true)

	// is case sensitive
	check("*BAZ*", "foo/bar/baz", false)
	check("*FOO*", "foo/bar/baz", false)
	check("A", "a", false)

	// Accepts array of patterns
	check([]string{"a*", "b*"}, "a/", true)
	check([]string{"a*", "b*"}, "b/", true)
	check([]string{"a*", "b*"}, "f/", false)

	check("*/foo", "blah/baz/foo", true)

	// empty patterns
	check([]string{}, "test", false)
	check([]string{}, "", true)
}

func TestIsJSON(t *testing.T) {
	check := func(value string, expected bool) {
		actual := IsJSON(value)
		if actual != expected {
			t.Errorf("Expected <%t> got <%t>", expected, actual)
		}
	}

	check(`1`, true)
	check(`"[1,2,3]"`, true)
	check(`"[1,   2,   3]"`, true)
	check(`{"first": "John", "last": "Doe"}`, true)
	check(`[{"first": "John", "last": "Doe"}, {"first": "Jane", "last": "Doe"}]`, true)
	check("1,", false)
	check("[1,2,3", false)
	check("[1,   2   3]", false)
	check(`{first: "John"}`, false)
	check(`[{first: "John"}, {first: "Jane"}]`, false)
	check("", false)
}

func TestIsUrl(t *testing.T) {

	check := func(value string, expected bool) {
		actual := IsUrl(value)
		if actual != expected {
			t.Errorf("Expected <%t> got <%t>", expected, actual)
		}
	}

	check(`https://go.dev`, true)
	check(`go.dev`, false)
	check(`http://localhost`, true)
	check(`invalid url`, false)
}

func TestIsUUID(t *testing.T) {

	check := func(value string, expected bool) {
		actual := IsUUID(value)
		if actual != expected {
			t.Errorf("Expected <%t> got <%t>", expected, actual)
		}
	}

	check(`fd2ac93c-6783-4fa9-830f-a32dff18fb0b`, true)
	check(`71ebdc83-9df4-4409-ad61-00d53707a9a3`, true)
	check(`d7115ae3-4291-4901-b5f7-343da54d8146`, true)
	check(`03a88d85-552a-4242-8193-85fb963ba529`, true)
	check(`b69ffe23-5a89-4d6f-99b5-65515b52c212`, true)
	check(`713d996e-e1db-4c7d-b94a-0aa1dd1b4be7`, true)
	check(`b487f049-2777-41cc-8164-0ed6ec30f6db`, true)
	check(`ec3feb6d-e116-48a8-a237-0dc0724821ed`, true)
	check(`3225d006-e67c-48d1-b4ca-acce30baca43`, true)
	check(`283d2f17-43f6-4c6d-939f-4f95ffd5b3f0`, true)

	check(`fd2ac93c-6783-830f-a32dff18fb0b`, false)
	check(`71e-9df4-4409-ad61-00d53707a9a3`, false)
	check(`d7115ae3-4291-4901-b5f7146`, false)
	check(`-85fb963ba529`, false)
	check(`12345`, false)
	check(`713d996ee1db4c7db94a-0aa1dd1b4be7`, false)

}

func TestIsULID(t *testing.T) {

	check := func(value string, expected bool) {
		actual := IsULID(value)
		if actual != expected {
			fmt.Println(value)
			t.Errorf("Expected <%t> got <%t>", expected, actual)
		}
	}

	check(`01GJSNW9MAF792C0XYY8RX6QFT`, true)
	check(`01GJSNW9MAF-792C0XYY8RX6ssssss-QFT`, false)
}

func TestSnake(t *testing.T) {
	check := func(value, expected string) {
		actual := Snake(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("Golang string_helpers", "golang_string_helpers")
	check("Golang_string_helpers", "golang_string_helpers")
	check("Golang string helpers", "golang_string_helpers")
	check("GolangStringHelpers", "golang_string_helpers")
	check("Golang-string-helpers", "golang_string_helpers")
	check("Foo-Bar", "foo_bar")
	check("Foo_Bar", "foo_bar")
	check("ŻółtaŁódka", "żółtałódka")
}

func TestKebab(t *testing.T) {
	check := func(value, expected string) {
		actual := Kebab(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("Golang string_helpers", "golang-string-helpers")
	check("Golang_string_helpers", "golang-string-helpers")
	check("Golang string helpers", "golang-string-helpers")
	check("GolangStringHelpers", "golang-string-helpers")
	check("Golang-string-helpers", "golang-string-helpers")
	check("Foo-Bar", "foo-bar")
	check("Foo_Bar", "foo-bar")
	check("ŻółtaŁódka", "żółtałódka")
}

func TestLcfirst(t *testing.T) {

	check := func(value, expected string) {
		actual := Lcfirst(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("Chris", "chris")
	check("Go", "go")
	check("CHRIS", "cHRIS")
	check("", "")
	check("A", "a")
}

func TestLength(t *testing.T) {
	check := func(value string, expected int) {
		actual := Length(value)
		if actual != expected {
			t.Errorf("Expected <%d> got <%d>", expected, actual)
		}
	}

	check("", 0)
	check("chris@example.com", 17)
	check("foo bar baz", 11)

	// Test for multibyte string support
	check("这是一段中文", 6)
	check("Jönköping", 9)
}

func TestLimit(t *testing.T) {
	check := func(value string, limit int, expected string) {
		actual := Limit(value, limit)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	check("chris", 3, "chr...")
	check(s, 10, "Lorem ipsu...")
	check(s, 20, "Lorem ipsum dolor si...")
	check(s, 30, "Lorem ipsum dolor sit amet, co...")
	check(s, 100, s)
	check(s, 0, s)

	nonAsciiString := "这是一段中文"

	check(nonAsciiString, 1, "这...")
	check(nonAsciiString, 3, "这是一...")
	check(nonAsciiString, 4, "这是一段...")
	check(nonAsciiString, 5, "这是一段中...")
	check(nonAsciiString, 100, nonAsciiString)
	check(nonAsciiString, 0, nonAsciiString)
}

func TestLower(t *testing.T) {

	check := func(value, expected string) {
		actual := Lower(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("FOO BAR BAZ", "foo bar baz")
	check("fOo Bar bAz", "foo bar baz")
}

func TestMask(t *testing.T) {

	check := func(value, character string, index, length int, expected string) {
		actual := Mask(value, character, index, length)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("abcdef", "*", 0, -1, "*****f")
	check("abcdef", "*", 2, -1, "ab***f")
	check("abcdef", "*", 4, -4, "abcdef")
	check("abcdef", "*", -3, -1, "abc**f")

	check("chris@example.com", "*", 0, 2, "**ris@example.com")
	check("chris@example.com", "*", 3, 0, "chris@example.com")
	check("chris@example.com", "*", 0, 1, "*hris@example.com")
	check("chris@example.com", "*", -6, 3, "chris@examp***com")

	check("chris@example.com", "*", -15, 4, "ch****example.com")
	check("chris@example.com", "-", -15, 4, "ch----example.com")
	check("chris@example.com", "$", -15, 4, "ch$$$$example.com")

	check("chris@example.com", "*", 3, 8, "chr********le.com")
	check("chris@example.com", "*", -17, 0, "chris@example.com")
	check("chris@example.com", "*", -177, 1, "*****************")
	check("chris@example.com", "*", 10, 10, "chris@exam*******")
	check("chris@example.com", "*", 10, 7, "chris@exam*******")
	check("chris@example.com", "*", 16, 99, "chris@example.co*")
	check("chris@example.com", "*", 17, 99, "chris@example.com")

	// Test for multibyte string support
	check("这是一段中文", "*", 0, 2, "**一段中文")
	check("这是一段中文", "*", -4, 2, "这是**中文")
	check("这是一段中文", "*", 0, 1, "*是一段中文")
	check("Jönköping", "*", 0, 1, "*önköping")
	check("Jönköping", "*", -6, 4, "Jön****ng")
}

func TestMatch(t *testing.T) {

	check := func(pattern, value, expected string) {
		actual := Match(pattern, value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check(`foo.?`, "seafood fool", "food")
	check(`foo.?`, "meat", "")
}

func TestMatchAll(t *testing.T) {

	check := func(pattern, value string, expected []string) {
		actual := MatchAll(pattern, value)

		if len(actual) == 0 && len(expected) == 0 {
			return
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected <%+v> got <%+v>", expected, actual)
		}
	}

	check(`foo.?`, "seafood fool", []string{"food", "fool"})
	check(`foo.?`, "meat", []string{})
	check(`a.`, "paranormal", []string{"ar", "an", "al"})
	check(`a.`, "none", []string{})
}

func TestNumbers(t *testing.T) {

	check := func(value, expected string) {
		actual := Numbers(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("chr15k", "15")
	check("abc123", "123")
	check("(555) 123-4567", "5551234567")
	check("chris", "")
	check("12345", "12345")
	check("1-2$3££/@**4***5", "12345")
}

func TestPadBoth(t *testing.T) {

	check := func(value string, length int, pad, expected string) {
		actual := PadBoth(value, length, pad)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("Chris", 14, "-", "----Chris-----")
	check("Chris", 10, "+=", "+=Chris+=+")
	check("Chris", 9, "+=+=+", "+=Chris+=")

	// Test for multibyte string support
	check("Chris", 16, "❤☆", "❤☆❤☆❤Chris❤☆❤☆❤☆")
	check("这是一段中文", 10, "_", "__这是一段中文__")
	check("Jönköping", 12, "*", "*Jönköping**")
}

func TestPadLeft(t *testing.T) {

	check := func(value string, length int, pad, expected string) {
		actual := PadLeft(value, length, pad)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("Chris", 6, "-", "-Chris")
	check("Chris", 10, "+=", "+=+=+Chris")
	check("Chris", 10, "+=+=+", "+=+=+Chris")

	// Test for multibyte string support
	check("Chris", 16, "❤☆", "❤☆❤☆❤☆❤☆❤☆❤Chris")
	check("这是一段中文", 10, "_", "____这是一段中文")
	check("Jönköping", 12, "*", "***Jönköping")
}

func TestPadRight(t *testing.T) {

	check := func(value string, length int, pad, expected string) {
		actual := PadRight(value, length, pad)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("Chris", 6, "-", "Chris-")
	check("Chris", 10, "+=", "Chris+=+=+")
	check("Chris", 10, "+=+=+", "Chris+=+=+")

	// Test for multibyte string support
	check("Chris", 16, "❤☆", "Chris❤☆❤☆❤☆❤☆❤☆❤")
	check("这是一段中文", 10, "_", "这是一段中文____")
	check("Jönköping", 12, "*", "Jönköping***")
}

func TestPassword(t *testing.T) {

	check := func(length int, includeNumber, includeSpecial bool) {
		actual := Password(length, includeNumber, includeSpecial)
		if length != len(actual) {
			t.Errorf("Expected length <%d> got <%d>", length, len(actual))
		}
	}

	check(0, false, false)
	check(8, true, true)
	check(16, false, true)
	check(32, true, false)
	check(2048, true, true)
}

func TestRandom(t *testing.T) {

	check := func(length int) {
		actual := Random(length)
		if length != len(actual) {
			t.Errorf("Expected length <%d> got <%d>", length, len(actual))
		}

		// alphanumeric check
		if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(actual) {
			t.Errorf("Expected alphanumeric chars only, got <%s>", actual)
		}
	}

	check(0)
	check(8)
	check(16)
	check(32)
	check(256)
	check(2048)
}

func TestSlug(t *testing.T) {

	check := func(value string, dictionary map[string]string, expected string) {
		actual := Slug(value, dictionary)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	overrides := make(map[string]string)
	overrides["%"] = "-percent-"
	overrides["!"] = "-bang-"
	overrides["xxx"] = "-123-"
	overrides["$"] = "-dollar-"
	overrides["£"] = "-pound-"

	check("500-$-bill!", overrides, "500-dollar-bill-bang")
	check("user@host", overrides, "user-at-host")
	check("hello world", overrides, "hello-world")
	check("hello-world?", overrides, "hello-world")
	check("hello_world!", overrides, "hello-world-bang")
	check("user@host", overrides, "user-at-host")
	check("xxx", overrides, "123")
	check("500$ bill", overrides, "500-dollar-bill")
	check("500--$----bill", overrides, "500-dollar-bill")
	check("500-$-bill", overrides, "500-dollar-bill")
	check("500$--bill", overrides, "500-dollar-bill")
	check("500-$--bill", overrides, "500-dollar-bill")
	check("500-$-bill!", overrides, "500-dollar-bill-bang")
	check("500-£-bill-xxx-%", overrides, "500-pound-bill-123-percent")
}

func TestSquish(t *testing.T) {

	check := func(value, expected string) {
		actual := Squish(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("foo bar   baz  ", "foo bar baz")
	check("   foo     bar  baz  ", "foo bar baz")
}

func TestStartsWith(t *testing.T) {

	check := func(haystack string, needles interface{}, expected bool) {
		actual := StartsWith(haystack, needles)
		if actual != expected {
			t.Errorf("Expected <%t> got <%t>", expected, actual)
		}
	}

	check("jason", "jas", true)
	check("jason", "jason", true)
	check("jason", []string{"jas"}, true)
	check("jason", []string{"day", "jas"}, true)
	check("jason", "day", false)
	check("jason", []string{"day"}, false)
	check("jason", nil, false)
	check("jason", []string{}, false)
	check("0123", []string{}, false)
	check("0123", "0", true)
	check("jason", "J", false)
	check("jason", "", false)
	check("", "", false)
	check("7", " 7", false)
	check("7a", "7", true)
	// Test for multibyte string support
	check("Jönköping", "Jö", true)
	check("Malmö", "Malmö", true)
	check("Jönköping", "Jonko", false)
	check("Malmö", "Malmo", false)
	check("你好", "你", true)
	check("你好", "好", false)
	check("你好", "a", false)
}

func TestStudly(t *testing.T) {

	check := func(value, expected string) {
		actual := Studly(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("foo bar baz", "FooBarBaz")
	check("foo-bar-baz", "FooBarBaz")
	check("foo_bar_baz", "FooBarBaz")
	check("fooBarBaz", "FooBarBaz")
}

func TestTake(t *testing.T) {

	check := func(value string, limit int, expected string) {
		actual := Take(value, limit)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("abcdef", 2, "ab")
	check("abcdef", -2, "ef")
	check("abcdef", 0, "")
	check("", 2, "")
	check("abcdef", 10, "abcdef")
	check("abcdef", 6, "abcdef")
	check("üöä", 1, "ü")
}

func TestTrim(t *testing.T) {

	check := func(value, expected string) {
		actual := Trim(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check(" Chris ", "Chris")
}

func TestTrimLeft(t *testing.T) {

	check := func(value, expected string) {
		actual := TrimLeft(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check(" Chris ", "Chris ")
}

func TestTrimRight(t *testing.T) {

	check := func(value, expected string) {
		actual := TrimRight(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check(" Chris ", " Chris")
}

func TestUcfirst(t *testing.T) {

	check := func(value, expected string) {
		actual := Ucfirst(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("chris", "Chris")
	check("go", "Go")
	check("cHRis", "CHRis")
	check("", "")
	check("a", "A")
}

func TestUpper(t *testing.T) {
	check := func(value, expected string) {
		actual := Upper(value)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("foo bar baz", "FOO BAR BAZ")
	check("fOo Bar bAz", "FOO BAR BAZ")
}

func TestUnwrap(t *testing.T) {
	check := func(value, before, after, expected string) {
		actual := Unwrap(value, before, after)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("`value`", "`", "", "value")
	check("`value", "`", "", "value")
	check("value`", "`", "", "value")
	check("foo-bar-baz", "foo-", "-baz", "bar")
	check("{some: 'json'}", "{", "}", "some: 'json'")
}

func TestWords(t *testing.T) {

	check := func(value string, words int, expected string) {
		actual := Words(value, words)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	check(s, 2, "Lorem ipsum...")
	check(s, 4, "Lorem ipsum dolor sit...")
	check(s, 5, "Lorem ipsum dolor sit amet,...")
	check(s, 6, "Lorem ipsum dolor sit amet, consectetur...")
	check(s, 100, s)
	check(s, 0, "...")
	check("foo bar baz", 2, "foo bar...")
	check("foo bar baz", 3, "foo bar baz")
}

func TestWrap(t *testing.T) {

	check := func(value, before, after, expected string) {
		actual := Wrap(value, before, after)
		if actual != expected {
			t.Errorf("Expected <%s> got <%s>", expected, actual)
		}
	}

	check("value", "`", "", "`value`")
	check("-bar-", "foo", "baz", "foo-bar-baz")
}
