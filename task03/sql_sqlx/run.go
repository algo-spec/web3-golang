package sql_sqlx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

/**
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
*/

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float32 `db:"salary"`
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

func Run(db *sqlx.DB) {
	var employees []Employee
	err := db.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department = '技术部'")
	if err != nil {
		panic(err)
	}
	fmt.Println("查询技术部员工信息：", employees)

	var highestSalaryEmployee Employee
	err = db.Get(&highestSalaryEmployee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println("工资最高的员工信息：", highestSalaryEmployee)

	var books []Book
	err = db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > 50")
	if err != nil {
		panic(err)
	}
	fmt.Println("查询价格大于50元的书籍：", books)
}
