package mysql

import (
	"database/sql"

	"github.com/rmftelier/Asulside/pkg/models"

)

type BlogModel struct{
     DB *sql.DB
}

func(m *BlogModel)Insert(title, article string) (int, error){

  stmt := `INSERT INTO blogs (title, article, publishedAt)

    VALUES(?, ?, UTC_TIMESTAMP())`

  result, err := m.DB.Exec(stmt, title, article)
  if err != nil{
     return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil{
      return 0, err
  }

  return int(id), nil
} 

func(m *BlogModel) Get(id int)(*models.Blog, error){
    stmt := `SELECT id, title, article, publishedAt FROM blogs
  
    WHERE id = ?`

    row := m.DB.QueryRow(stmt, id)

    b := &models.Blog{}

    err := row.Scan(&b.ID, &b.Title, &b.Article, &b.PublishedAt)

    if err == sql.ErrNoRows{
        return nil, models.ErrNoRecord
    }else if err != nil{
       return nil, err
    }

    return b, nil
}

func(m *BlogModel) Latest()([]*models.Blog, error){
    
  stmt := `SELECT id, title, article, publishedAt FROM blogs 
    ORDER BY publishedAt DESC LIMIT 10`

    rows, err := m.DB.Query(stmt)
    if err != nil{
        return nil, err
    }
    defer rows.Close()

    blogs := []*models.Blog{}
    for rows.Next(){
        b := &models.Blog{}
        err = rows.Scan(&b.ID, &b.Title, &b.Article, &b.PublishedAt)
        if err != nil{
            return nil, err
        }
        blogs = append(blogs, b)
    }

    err = rows.Err()
    if err != nil{
        return nil, err
    }
  
  return blogs, nil
}