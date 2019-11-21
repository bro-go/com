package bstr

import "testing"

type User struct {
	Name     string
	Age      int
	Password string
	Next     *User
}

func TestDD(t *testing.T) {
	DD("Hello")

	DD([]int{1, 2, 3, 4, 5})

	DD([]string{"Hello", "World", "我是一个字符串 slice"})

	user := &User{
		Name:     "张三",
		Age:      18,
		Password: "82390281093812djskla",
		Next: &User{
			Name:     "杨炯突然",
			Age:      111,
			Password: "ksjdlksajdlksjalkjdksaljkdlsadjkslajdl",
		},
	}

	DD(user)

	DD(map[string]string{
		"name": "张三",
		"年龄":   "188",
		"性别":   "男",
	})

	DD(map[int]interface{}{
		1: "hello",
		2: []int{1, 2, 3, 4, 5},
		3: []string{"heelo", "HHHHH", "abc"},
		4: 12.34556,
	})
}
