package dao

import (
	"Library/Library/model"
	"Library/Library/utils"
)

func AddComment(comment *model.Comment) error {
	sqlStr := "insert into comment(id,book_id,user_id,time,text,username) values(?,?,?,?,?,?)"
	utils.Db.Exec(sqlStr, comment.ID, comment.BookID, comment.UserID, comment.Time, comment.Text, comment.UserName)
	return nil
}
func GetAllComment(bookID int) ([]*model.Comment, error) {
	sqlStr := "select comment.id,book_id,user_id,time,text,comment.username from comment  join user u  on comment.user_id = u.id where book_id = ?"
	rows, _ := utils.Db.Query(sqlStr, bookID)
	var comments []*model.Comment
	for rows.Next() {
		comment := &model.Comment{}
		err := rows.Scan(&comment.ID, &comment.BookID, &comment.UserID, &comment.Time, &comment.Text, &comment.UserName)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
func DeleteComment(commentID int) error {
	sqlStr := "delete from comment where id=?"
	_, err := utils.Db.Exec(sqlStr, commentID)
	if err != nil {
		return err
	}
	return nil
}
