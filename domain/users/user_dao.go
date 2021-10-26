package users

import (
	"fmt"

	"github.com/JizongL/bookstore_users-api/datasources/mysql/users_db"
	"github.com/JizongL/bookstore_users-api/utils/date_utils"
	"github.com/JizongL/bookstore_users-api/utils/errors"
	"github.com/JizongL/bookstore_users-api/utils/mysql_utils"
)

const(
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name,last_name,email, date_created) VALUES(?,?,?,?);"
	queryGetUser = "SELECT id, first_name, last_name,email,date_created FROM users WHERE id=?"
)

var(
	usersDB = make(map[int64]*User)
)

func (user *User)Get()*errors.RestErr{
	stmt,err:= users_db.Client.Prepare(queryGetUser)
	if err!=nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result:=stmt.QueryRow(user.Id)
	if getErr:=result.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated);getErr!=nil{	
		return mysql_utils.ParseError(getErr)
		// if strings.Contains(err.Error(),errorNoRows){
		// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found",user.Id))}		
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d %s",user.Id,getErr.Error()))
	}
	// if err:= users_db.Client.Ping();err!=nil{
	// 	panic(err)
	// }
	// result:= usersDB[user.Id]
	// if result==nil{
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found",user.Id))
	// }
	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated
	return nil
}

func (user *User)Save()*errors.RestErr{
	stmt,err:= users_db.Client.Prepare(queryInsertUser)
	if err!=nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult,saveErr:= stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated)
	if saveErr != nil{
		return mysql_utils.ParseError(saveErr)
		
		// type assertion technique used here. 
		// sqlErr,ok:= saveErr.(*mysql.MySQLError)
		// if !ok{
		// 	return errors.NewBadRequestError(fmt.Sprintf("error when trying to save user %s",sqlErr))
		// }
		// switch sqlErr.Number{
		// 	case 1062:
		// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already exists",user.Email))
		// }
		// return errors.NewBadRequestError(fmt.Sprintf("error when trying to save user %s",saveErr))
	}
	// if err!=nil{
	// 	if strings.Contains(err.Error(),indexUniqueEmail){
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already exists",user.Email))
	// 	}
	// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s",err.Error()))
	// }
	// fmt.Println(insertResult)
	userId,err:= insertResult.LastInsertId()
	if err!=nil{
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s",err.Error()))
	}
	user.Id = userId
	// the following code was for in memory db setup at the begining. 
	// current:= usersDB[user.Id]
	// if current!=nil{
	// 	if current.Email==user.Email{
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered",user.Email))	
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists",user.Id))
	// }
	// user.DateCreated=date_utils.GetNowString()
	// usersDB[user.Id]=user
	return nil
}