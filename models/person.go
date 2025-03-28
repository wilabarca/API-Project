package models



type Person struct {
    ID     int    `json:"id"`
    Nombre string `json:"nombre"`
    Edad   int    `json:"edad"`
    Genero string `json:"genero"`
    Sexo   string `json:"sexo"`
}

func (p *Person) Create() error {
    query := `INSERT INTO persons (nombre, edad, genero, sexo) VALUES (?, ?, ?, ?)`
    result, err := DB.Exec(query, p.Nombre, p.Edad, p.Genero, p.Sexo)
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }

    p.ID = int(id)
    return nil
}

func GetAllPersons() ([]Person, error) {
    query := `SELECT id, nombre, edad, genero, sexo FROM persons`
    rows, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var persons []Person
    for rows.Next() {
        var p Person
        if err := rows.Scan(&p.ID, &p.Nombre, &p.Edad, &p.Genero, &p.Sexo); err != nil {
            return nil, err
        }
        persons = append(persons, p)
    }

    return persons, nil
}

func GetGenderCounts() (map[string]int, error) {
    query := `SELECT genero, COUNT(*) as count FROM persons GROUP BY genero`
    rows, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    counts := make(map[string]int)
    for rows.Next() {
        var genero string
        var count int
        if err := rows.Scan(&genero, &count); err != nil {
            return nil, err
        }
        counts[genero] = count
    }

    return counts, nil
}