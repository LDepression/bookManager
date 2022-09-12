package dao

import (
	"Library/Library/model"
	"Library/Library/utils"
	"net/http"
)

func AddSession(session *model.Session) error {
	defer utils.RwMutexSession.Unlock()
	utils.RwMutexSession.Lock()
	sqlStr := "insert into session(session_id,username,user_id) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, session.SessionID, session.UserName, session.UserID)
	if err != nil {
		return err
	}
	return nil
}
func DeleteSession(session *model.Session) error {
	defer utils.RwMutexSession.Unlock()
	utils.RwMutexSession.Lock()
	sqlStr := "delete from session where session_id=?"
	_, err := utils.Db.Exec(sqlStr, session.SessionID)
	if err != nil {
		return err
	}
	return nil
}

//GetSessionByID 根据sessionID的值从数据库中查询session
func GetSessionByID(cookieID string) *model.Session {
	defer utils.RwMutexSession.RUnlock()
	utils.RwMutexSession.RLock()
	sqlStr := "select session_id,username,user_id from session where session_id=?"
	row := utils.Db.QueryRow(sqlStr, cookieID)
	session := &model.Session{}
	row.Scan(&session.SessionID, &session.UserName, &session.UserID)
	return session
}
func IsLogin(r *http.Request) (*model.Session, bool) {
	//先获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//getSessionByID
		cookieValue := cookie.Value
		session := GetSessionByID(cookieValue)
		if session.UserID > 0 {
			return session, true
		}
	}
	return nil, false
}
