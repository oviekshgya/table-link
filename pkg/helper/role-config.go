package helper

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
	"sync"
	"table-link/src/model/role"
	"time"
)

var (
	lastCheckedTime time.Time = time.Now()
	Mu              sync.Mutex
)
var ListRoleData []*Role

type Role struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique;not null"`
	RoleRights []RoleRight
}

type RoleRight struct {
	ID      uint   `gorm:"primaryKey"`
	RoleID  uint   `gorm:"not null"`
	Section string `gorm:"not null"`
	Route   string `gorm:"not null"`
	RCreate bool   `gorm:"default:false"`
	RRead   bool   `gorm:"default:false"`
	RUpdate bool   `gorm:"default:false"`
}

func GetRoleData(db *gorm.DB) {
	fmt.Println("GetRoleData starting")
	var data []*Role
	var clinet = role.RoleModel{
		DB: db,
	}

	result, err := clinet.GetRoleWithRights()
	if err != nil {
		log.Println("Role Get Rights error", err)
	}
	if len(result) > 0 {
		for i, val := range result {
			data = append(data, &Role{
				ID:   val.ID,
				Name: val.Name,
				RoleRights: []RoleRight{
					RoleRight{
						ID:      val.RoleRights[i].ID,
						RoleID:  val.RoleRights[i].RoleID,
						Section: val.RoleRights[i].Section,
						Route:   val.RoleRights[i].Route,
						RCreate: val.RoleRights[i].RCreate,
						RRead:   val.RoleRights[i].RRead,
						RUpdate: val.RoleRights[i].RUpdate,
					},
				},
			})
		}
		ListRoleData = data
	} else {
		log.Println("FindAllClient error array")
	}

}

func StartAddClient(ctx context.Context, db *gorm.DB) {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	clientModel := role.RoleModel{DB: db}

	for {
		select {
		case <-ticker.C:
			_, cancel := context.WithTimeout(ctx, 10*time.Second)
			activeClients, err := clientModel.GetRoleWithRights()
			cancel()

			if err != nil {
				log.Println("Role error:", err)
				continue
			}

			syncClientList(activeClients)

		case <-ctx.Done():
			log.Println("Shutting down Role...")
			return
		}
	}
}

func syncClientList(activeClients []*role.Role) {
	Mu.Lock()
	defer Mu.Unlock()

	activeIPMap := make(map[string]bool)
	for _, client := range activeClients {
		activeIPMap[client.Name] = true
	}

	var updatedList []*Role
	for _, client := range ListRoleData {
		if activeIPMap[client.Name] {
			updatedList = append(updatedList, client)
		} else {
			log.Println("Removing inactive Role:", client.Name)
		}
	}

	for _, client := range activeClients {
		exists := false
		for _, existing := range updatedList {
			if existing.Name == client.Name {
				exists = true
				break
			}
		}
		if !exists {
			log.Println("Adding new active Role:", client.Name)
			updatedList = append(updatedList, &Role{Name: client.Name})
		}
	}

	ListRoleData = updatedList
	var newList []string
	if len(updatedList) > 0 {
		for _, val := range updatedList {
			newList = append(newList, val.Name)
		}
	}
	log.Println("Updated Role: ", strings.Join(newList, ", "))
}
