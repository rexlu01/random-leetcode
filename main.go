package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"time"

	"gopkg.in/yaml.v2"
)

//YamlOf is xxx
type YamlOf struct {
	HardCode []int    `yaml:"HardCode"`
	Type     []string `yaml:"Type"`
	CodeMap  `yaml:"CodeMap"`
}

//CodeMap is xxx
type CodeMap struct {
	Array              []int `yaml:"Array"`
	String             []int `yaml:"String"`
	DynamicProgramming []int `yaml:"DynamicProgramming"`
	HashTable          []int `yaml:"HashTable"`
	Math               []int `yaml:"Math"`
	DepthFirstSearch   []int `yaml:"DepthFirstSearch"`
	Sorting            []int `yaml:"Sorting"`
	Greedy             []int `yaml:"Greedy"`
	BreadthFirstSearch []int `yaml:"BreadthFirstSearch"`
	Tree               []int `yaml:"Tree"`
	BinarySearch       []int `yaml:"BinarySearch"`
	BinaryTree         []int `yaml:"BinaryTree"`
	Matrix             []int `yaml:"Matrix"`
	TwoPointers        []int `yaml:"TwoPointers"`
	BitManipulation    []int `yaml:"BitManipulation"`
	Stack              []int `yaml:"Stack"`
	Design             []int `yaml:"Design"`
	Backtracking       []int `yaml:"Backtracking"`
	Graph              []int `yaml:"Graph"`
	Simulation         []int `yaml:"Simulation"`
	SlidingWindow      []int `yaml:"SlidingWindow"`
	Counting           []int `yaml:"Counting"`
	LinkedList         []int `yaml:"LinkedList"`
	UnionFind          []int `yaml:"UnionFind"`
	Recursion          []int `yaml:"Recursion"`
	BinarySearchTree   []int `yaml:"BinarySearchTree"`
	Trie               []int `yaml:"Trie"`
	DivideandConquer   []int `yaml:"DivideandConquer"`
	Memoization        []int `yaml:"Memoization"`
	HashFunction       []int `yaml:"HashFunction"`
	StringMatching     []int `yaml:"StringMatching"`
	RollingHash        []int `yaml:"RollingHash"`
	Combinatorics      []int `yaml:"Combinatorics"`
	Iterator           []int `yaml:"Iterator"`
	MergeSort          []int `yaml:"MergeSort"`
	DoublyLinkedList   []int `yaml:"DoublyLinkedList"`
}

//TempCode is xxx
type TempCode struct {
	Day             int    `yaml:"Day"`
	TypeName        string `yaml:"TypeName"`
	CurrentCodeList []int  `yaml:"CurrentCodeList"`
}

func main() {
	day := time.Now().Day()
	te := ReadYamlTemp()
	ya := ReadYamlConf()

	if te.Day != day {
		te.Day = day
		WriteYamlTemp(te)
		//选随机type
		typeList := ya.Type
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		typeNum := r.Intn(len(typeList))
		typeName := typeList[typeNum]
		te.TypeName = typeName
		//通过typeName获取
		CodeList := GetCodeList(typeName)
		//随机CodeList中的一题
		i := r.Intn(len(CodeList))
		Code := CodeList[i]
		if len(CodeList) != 0 {
			CodeList = append(CodeList[:i], CodeList[i+1:]...)
		} else {
			CodeList = []int{}
		}
		te.CurrentCodeList = CodeList
		WriteYamlTemp(te)
		fmt.Printf("题目类型 %s, 题目列表 %v", typeName, Code)
		fmt.Println()
	} else {
		if len(te.CurrentCodeList) == 0 {
			typeList := ya.Type
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			typeNum := r.Intn(len(typeList))
			typeName := typeList[typeNum]
			te.TypeName = typeName
			//通过typeName获取
			CodeList := GetCodeList(typeName)
			//随机CodeList中的一题
			i := r.Intn(len(CodeList))
			Code := CodeList[i]
			if len(CodeList) != 0 {
				CodeList = append(CodeList[:i], CodeList[i+1:]...)
			} else {
				CodeList = []int{}
			}
			te.CurrentCodeList = CodeList
			WriteYamlTemp(te)
			fmt.Printf("题目类型 %s, 题目列表 %v", typeName, Code)
			fmt.Println()
		} else {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			CodeList := te.CurrentCodeList
			i := r.Intn(len(CodeList))
			Code := CodeList[i]
			if len(CodeList) != 0 {
				CodeList = append(CodeList[:i], CodeList[i+1:]...)
			} else {
				CodeList = []int{}
			}
			te.CurrentCodeList = CodeList
			WriteYamlTemp(te)
			fmt.Printf("题目类型 %s, 题目列表 %v", te.TypeName, Code)
			fmt.Println()
		}

	}

}

//GetCodeList is xxx
func GetCodeList(typeName string) []int {
	ya := new(YamlOf)
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Printf("%#v", err)
	}
	err = yaml.Unmarshal(yamlFile, ya)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	t := reflect.TypeOf(ya.CodeMap)

	field, Flag := t.FieldByName(typeName)
	if Flag {
		CodeMapIndex := field.Index[0]
		v := reflect.ValueOf(ya.CodeMap)
		t := v.Field(CodeMapIndex).Interface()
		//fmt.Printf("%T", t.([]string))
		return t.([]int)
	}
	return []int{}
}

//ReadYamlTemp is xxx
func ReadYamlTemp() (te *TempCode) {
	//写yaml
	//初始化
	te = new(TempCode)

	//打开temp.yaml
	yamlFile, err := ioutil.ReadFile("temp.yaml")
	if err != nil {
		fmt.Printf("%#v", err)
	}
	err = yaml.Unmarshal(yamlFile, te)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	return te

}

//WriteYamlTemp is xxx
func WriteYamlTemp(te *TempCode) {
	outTe, err := yaml.Marshal(te)
	if err != nil {
		fmt.Printf("%#v", err)
	}
	ioutil.WriteFile("temp.yaml", outTe, 0777)

}

//ReadYamlConf is xxx
func ReadYamlConf() (ya *YamlOf) {
	ya = new(YamlOf)
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Printf("%#v", err)
	}
	err = yaml.Unmarshal(yamlFile, ya)
	if err != nil {
		fmt.Printf("%#v", err)
		return ya
	}
	return ya
}
