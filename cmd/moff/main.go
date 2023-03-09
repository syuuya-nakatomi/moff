package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Result struct {
	Packages []Package `json:"packages"`
}

func main() {
	// 引数からファイル名を取得する
	filename := os.Args[1]

	// JSONファイルを読み込む
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// JSONデータを構造体に変換する
	var result Result
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	// Ansible playbookを生成する
	playbook := "---\n- name: Install packages\n  hosts: all\n  become: true\n  tasks:\n    - name: Install package\n      apt:\n        name: \"{{ item.name }}={{ item.version }}\"\n        state: present\n      with_items:\n"
	for _, pkg := range result.Packages {
		playbook += fmt.Sprintf("        - { name: \"%s\", version: \"%s\" }\n", pkg.Name, pkg.Version)
	}

	// Ansible playbookをファイルに書き込む
	err = ioutil.WriteFile("playbook.yml", []byte(playbook), 0644)
	if err != nil {
		panic(err)
	}

	// 成功した旨を表示する
	fmt.Println("Successfully created playbook.yml")
}
