
	// dal.Init()
	// var row model.User
	// mysql.DB.Create(&model.User{Email:"test@qq.com", PasswordHashed:"123"})//增加
	// mysql.DB.Model(&model.User{}).Where("Email = ?","test@qq.com").Update("PasswordHashed","22222")//改
	// mysql.DB.Model(&model.User{}).Where("Email = ?","test@qq.com").First(&row)//查
	// fmt.Println(row)
	// mysql.DB.Where("Email = ?","test@qq.com").Delete(&model.User{})//软删除
	// mysql.DB.Unscoped().Where("Email = ?","test@qq.com").Delete(&model.User{})//硬删除
	// mysql.DB.Migrator().DropTable("users")//删表
