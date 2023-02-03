package main

import (
	"testing"
)

func TestMonster_ReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("错误，希望为%v,实际为%v", true, res)

	}
	if monster.Name != "wqwd" {
		t.Fatalf("错误，希望为%v,实际为%v", "wqwd", monster.Name)

	}
	t.Logf("成功")
}

func TestMonster_Store(t *testing.T) {
	str := Monster{Name: "wqwd", Age: 12, Skill: "wiqdwd"}
	are := str.Store()
	if !are {
		t.Fatalf("测试错误，期望为%v,实际为%v", true, are)

	}
	t.Logf("测试成功")
}
