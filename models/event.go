package models

import (
	"fmt"
	"time"

	"example.com/api/db"
)
type Event struct {
    ID          int64 
    Name        string 
    Description string   
    Location    string 
    DateTime    time.Time
    UserID      string
} 
func (e *Event) Save() error {
    query := `INSERT INTO EVENTS(name, description, location, dateTime, userId)
        VALUES (?, ?, ?, ?, ?)`
    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()
    
    result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    e.ID = id
    return nil
}

func GetAllEvents() ([]Event, error) {
    query := `SELECT id, name, description, location, dateTime, userId FROM events`  
    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close() 

    var events []Event
    for rows.Next() {
        var event Event
        err := rows.Scan(&event.ID, &event.Name, &event.Description, 
            &event.Location, &event.DateTime, &event.UserID)
        if err != nil {
            
            return nil, err
        }
        events = append(events, event)
    }    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return events, nil
}

func GetEventById(id int64) (*Event,error) {
var event Event
query:= `SELECT * FROM events WHERE id = ?`
row:=  db.DB.QueryRow(query,id)
err:=row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)

if err != nil {
	return nil,err
}
	
return &event,nil
}


func (e Event) Update() error {
    query := `UPDATE events 
        SET name = ?, description = ?, location = ?, dateTime = ?, userId = ?
        WHERE id = ?`
    
    fmt.Printf("Updating event with ID: %d\n", e.ID)  // Debug log
    fmt.Printf("Values: %+v\n", e)  // Debug log
    
    stmt, err := db.DB.Prepare(query)
    if err != nil {
        fmt.Printf("Prepare error: %v\n", err)  // Debug log
        return err
    }
    defer stmt.Close()
    
    result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
    if err != nil {
        fmt.Printf("Exec error: %v\n", err)  // Debug log
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("no event found with ID %d", e.ID)
    }
    
    return nil
}

func (e Event) Delete() error {
query := `DELETE FROM events WHERE id = ?`
stmt,err := db.DB.Prepare(query)
if err != nil {
    fmt.Print(err)
    return err
}
defer stmt.Close()
_,err = stmt.Exec(e.ID)
if err != nil {
    fmt.Print(err)
    return err 
}
return nil 
}