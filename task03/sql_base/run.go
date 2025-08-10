package base_sql

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务
*/
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

type Account struct {
	ID      int
	Name    string
	Balance int
}

type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        int
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Student{})

	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	// 插入新记录
	db.Create(&student)

	fmt.Println("====================================================================")

	// 查询年龄大于18岁的学生
	var students []Student
	db.Debug().Where("age > ?", 18).Find(&students)
	fmt.Println(students)
	fmt.Println("====================================================================")

	// 更新姓名为"张三"的学生年级为"四年级"
	db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	fmt.Println("====================================================================")

	// 删除年龄小于15岁的学生记录
	db.Debug().Where("age < ?", 15).Delete(&Student{})
	fmt.Println("====================================================================")

	// 事务操作
	db.AutoMigrate(&Account{}, &Transaction{})
	accountA := Account{Name: "A", Balance: 1000}
	accountB := Account{Name: "B", Balance: 1000}
	db.Create(&accountA)
	db.Create(&accountB)

	err := db.Transaction(func(tx *gorm.DB) error {
		var accountA, accountB Account

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&Account{}).Where("name = ?", "A").First(&accountA).Error; err != nil {
			fmt.Println("查询A账户失败", err)
			return err
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&Account{}).Where("name = ?", "B").First(&accountB).Error; err != nil {
			fmt.Println("查询B账户失败", err)
			return err
		}

		if accountA.Balance < 100 {
			return gorm.ErrInvalidTransaction // 余额不足，回滚事务
		}
		fmt.Println("AccountA:", accountA)
		fmt.Println("AccountB:", accountB)

		accountA.Balance -= 100
		accountB.Balance += 100
		if err := tx.Save(accountA).Error; err != nil {
			fmt.Println("更新A账户失败", err)
			return err
		}

		if err := tx.Save(accountB).Error; err != nil {
			fmt.Println("更新B账户失败", err)
			return err
		}

		tx.Create(&Transaction{FromAccountID: accountA.ID, ToAccountID: accountB.ID, Amount: 100})
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
