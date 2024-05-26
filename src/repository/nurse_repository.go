package repository

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"health-record/model/database"
	"health-record/model/dto"
	"strconv"
	"time"

	"github.com/docker/distribution/uuid"
	"golang.org/x/crypto/bcrypt"
)

type NurseRepository struct {
	db *sql.DB
}

func NewNurseRepository(db *sql.DB) NurseRepositoryInterface {
	return &NurseRepository{db}
}

// CreateNurse inserts a new nurse into the database.

func (repo *NurseRepository) CreateNurse(ctx context.Context, nurse dto.RequestCreateNurse) (string, error) {
	// Generate a new password for the nurse
	password, err := GeneratePassword(12) // For example, generate a 12 character long password
	if err != nil {
			return "", err
	}

	// Hash the generated password before storing it
	hashedPassword, err := HashPassword(password)
	if err != nil {
			return "", err
	}

	// Prepare the SQL query to insert the new nurse with the hashed password
	const query = `INSERT INTO users (user_id, nip, name, role, identity_card_scan_img, password, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id`
	var userId string
	err = repo.db.QueryRowContext(ctx, query, uuid.Generate().String(), nurse.Nip, nurse.Name, "nurse", nurse.IdentityCardScanImg, hashedPassword, time.Now()).Scan(&userId)
	if err != nil {
			return "", err
	}

	// Optionally, you might want to handle sending the password to the nurse or displaying it as needed
	// For security reasons, do not log or display the raw password
	return userId, nil
}

func (ur *NurseRepository) GetUsers(ctx context.Context, param dto.RequestGetUser) ([]dto.UserDTO, error) {
	query := "SELECT user_id, nip, name, created_at FROM users WHERE 1=1"

	var args []interface{}

	if param.UserId != "" {
		query += " AND user_id LIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+param.UserId+"%")
	}

	if param.Name != "" {
		query += " AND name LIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+param.Name+"%")
	}

	if param.NIP != "" {
		query += " AND CAST(nip AS TEXT) LIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, param.NIP+"%")
	}

	if param.Role == "it" || param.Role == "nurse"{
		query += " AND role=$" + strconv.Itoa(len(args)+1)
		args = append(args, param.Role)
	}

	if param.CreatedAt == "asc" || param.CreatedAt == "desc" {
		query += fmt.Sprintf(" ORDER BY created_at %s", param.CreatedAt)
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", param.Limit, param.Offset)

	rows, err := ur.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dto.UserDTO
	for rows.Next() {
		var user dto.UserDTO
		if err := rows.Scan(&user.UserId, &user.NIP, &user.Name, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}


// UpdateNurse updates an existing nurse's information in the database.
func (repo *NurseRepository) UpdateNurse(ctx context.Context, userId string, nurse dto.RequestUpdateNurse) int {
	const query = `UPDATE users SET nip = $1, name = $2 WHERE user_id = $3`
	result, err := repo.db.ExecContext(ctx, query, nurse.Nip, nurse.Name, userId)
	fmt.Println("result>>>>>>", result)
	if err != nil {
		fmt.Printf("failed to update nurse: %v", err)
		return 500
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected>>>>>>", rowsAffected)
	if err != nil {
		fmt.Printf("failed to get rows affected: %v", err)
		return 404
	}
	if rowsAffected == 0 {
		fmt.Printf("nurse with user_id %s not found", userId)
		return 404
	}
	return 200
}

// DeleteNurse removes a nurse from the database.
func (r *NurseRepository) DeleteNurse(ctx context.Context, userId string) int {
	const query = `DELETE FROM users WHERE user_id = $1`
	// Execute the SQL query
	result, err := r.db.ExecContext(ctx, query, userId)
	fmt.Println("result>>>>>>", result)
	if err != nil {
		fmt.Printf("failed to delete nurse: %v", err)
		return 500
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected>>>>>>", rowsAffected)
	if err != nil {
		fmt.Printf("failed to get rows affected: %v", err)
		return 404
	}
	if rowsAffected == 0 {
		fmt.Printf("nurse with user_id %s not found", userId)
		return 404
	}

	return 200
}

func (r *NurseRepository) AddAccess(ctx context.Context, userId string, nurse dto.RequestAddAccess) int {
	const query = `UPDATE users SET password = $1 WHERE user_id = $2`

	// Hash the generated password before storing it
	hashedPassword, err := HashPassword(nurse.Password)
	if err != nil {
			return 500
	}

	result, err := r.db.ExecContext(ctx, query, hashedPassword, userId)
	fmt.Println("result>>>>>>", result)
	if err != nil {
		fmt.Printf("failed to update nurse: %v", err)
		return 500
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected>>>>>>", rowsAffected)
	if err != nil {
		fmt.Printf("failed to get rows affected: %v", err)
		return 404
	}
	if rowsAffected == 0 {
		fmt.Printf("nurse with user_id %s not found", userId)
		return 404
	}
	return 200
}

func GeneratePassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
			return "", err
	}
	password := base64.URLEncoding.EncodeToString(bytes)
	return password[:length], nil // Trimming the password in case base64 encoding exceeds the desired length
}

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
			return "", err
	}
	return string(hashedPassword), nil
}

func (r *NurseRepository) GetNurseByNIP(ctx context.Context, nip int64) (response database.User, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT user_id, name, nip, password FROM users WHERE nip = $1", nip).Scan(&response.Id, &response.Name, &response.Nip, &response.Password)
	if err != nil {
		return database.User{}, err
	}
	return response, nil
}

func (r *NurseRepository) GetNurseByID(ctx context.Context, userId string) (response database.User, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT user_id, name, nip, password FROM users WHERE user_id = $1", userId).Scan(&response.Id, &response.Name, &response.Nip, &response.Password)
	if err != nil {
		return database.User{}, err
	}
	return response, nil
}