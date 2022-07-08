package repository

import (
	"database/sql"
	"fmt"

	"gitlab.com/charlestenorios/next_ride_backend/app/domain"
)

type CysclistRepositoryPsql struct {
	db *sql.DB
}

var validation domain.Cyclist

func NewCyclistRepositoryPsql(db *sql.DB) *CysclistRepositoryPsql {
	return &CysclistRepositoryPsql{db: db}
}

func (c *CysclistRepositoryPsql) Get(id string) (domain.Cyclist, error) {
	var cyclist domain.Cyclist
	stmt, err := c.db.Prepare("SELECT id_cyclist, id_group, cyclist_name, cpf, birth," +
		" email, blood_type, health_plan, contact_emergency, contact_fone, got_to_know," +
		" create_at, active, img, participant_type, pedaling, tours, travels" +
		" where id_cyclist =$1")
	if err != nil {
		fmt.Print(err)
		return domain.Cyclist{}, err
	}
	err = stmt.QueryRow(id).Scan(&cyclist.Id, &cyclist.IdGroup, &cyclist.Name, &cyclist.Cpf, &cyclist.Birth,
		&cyclist.Email, &cyclist.BloodType, &cyclist.HealthPlan, &cyclist.ContactEmergency, &cyclist.GotToKnow,
		&cyclist.CreateAt, &cyclist.Active, &cyclist.Img, &cyclist.ParticipantType, &cyclist.Pedaling,
		&cyclist.Tours, &cyclist.Travels)
	if err != nil {
		fmt.Println(string(err.Error()))
		return domain.Cyclist{}, err
	}
	return cyclist, nil

}

func (c *CysclistRepositoryPsql) GetByCpf(cpf string) (domain.Cyclist, error) {
	var cyclist domain.Cyclist
	stmt, err := c.db.Prepare("SELECT id_cyclist, id_group, cyclist_name, cpf, birth," +
		" email, blood_type, health_plan, contact_emergency, contact_fone, got_to_know," +
		" create_at, active, img, participant_type, pedaling, tours, travels" +
		" where cpf =$1")
	if err != nil {
		fmt.Print(err)
		return domain.Cyclist{}, err
	}
	err = stmt.QueryRow(cpf).Scan(&cyclist.Id, &cyclist.IdGroup, &cyclist.Name, &cyclist.Cpf, &cyclist.Birth,
		&cyclist.Email, &cyclist.BloodType, &cyclist.HealthPlan, &cyclist.ContactEmergency, &cyclist.GotToKnow,
		&cyclist.CreateAt, &cyclist.Active, &cyclist.Img, &cyclist.ParticipantType, &cyclist.Pedaling,
		&cyclist.Tours, &cyclist.Travels)
	if err != nil {
		fmt.Println(string(err.Error()))
		return domain.Cyclist{}, err
	}
	return cyclist, nil

}

func (c *CysclistRepositoryPsql) GetByName(name string) (domain.Cyclist, error) {
	var cyclist domain.Cyclist
	stmt, err := c.db.Prepare("SELECT id_cyclist, id_group, cyclist_name, cpf, birth," +
		" email, blood_type, health_plan, contact_emergency, contact_fone, got_to_know," +
		" create_at, active, img, participant_type, pedaling, tours, travels" +
		" where cyclist_name like $1")
	if err != nil {
		fmt.Print(err)
		return domain.Cyclist{}, err
	}
	err = stmt.QueryRow(name).Scan(&cyclist.Id, &cyclist.IdGroup, &cyclist.Name, &cyclist.Cpf, &cyclist.Birth,
		&cyclist.Email, &cyclist.BloodType, &cyclist.HealthPlan, &cyclist.ContactEmergency, &cyclist.GotToKnow,
		&cyclist.CreateAt, &cyclist.Active, &cyclist.Img, &cyclist.ParticipantType, &cyclist.Pedaling,
		&cyclist.Tours, &cyclist.Travels)
	if err != nil {
		fmt.Println(string(err.Error()))
		return domain.Cyclist{}, err
	}
	return cyclist, nil

}

func (c *CysclistRepositoryPsql) Create(cyclist domain.Cyclist) (domain.Cyclist, error) {
	err := validation.Prepere()
	if err != nil {
		return domain.Cyclist{}, err
	}

	stmt, err := c.db.Prepare("INSERT INTO cyclist (id_cyclist, id_group, cyclist_name, cpf, birth, email," +
		" blood_type, health_plan, contact_emergency, contact_fone, got_to_know, create_at, active, img, " +
		"	participant_type, pedaling, tours, travels)" +
		" VALUES($1, $2, $3, $4, $5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)")
	if err != nil {
		fmt.Print(err.Error())

		return domain.Cyclist{}, err

	}
	_, err = stmt.Exec(cyclist.Id, cyclist.IdGroup, cyclist.Name, cyclist.Cpf, cyclist.Birth,
		cyclist.Email, cyclist.BloodType, cyclist.HealthPlan, cyclist.ContactEmergency, cyclist.GotToKnow,
		cyclist.CreateAt, cyclist.Active, cyclist.Img, cyclist.ParticipantType, cyclist.Pedaling,
		cyclist.Tours, cyclist.Travels)
	if err != nil {
		return domain.Cyclist{}, err
	}
	err = stmt.Close()
	if err != nil {
		return domain.Cyclist{}, err
	}
	return cyclist, nil
}

func (c *CysclistRepositoryPsql) FindAll() ([]domain.Cyclist, error) {
	var cyclist domain.Cyclist
	var cyclists []domain.Cyclist
	stmt, err := c.db.Query("SELECT id_cyclist, id_group, cyclist_name, cpf, birth, email, blood_type, health_plan, contact_emergency, contact_fone," +
		" got_to_know, create_at, active, img, participant_type, pedaling, tours, travels" +
		"FROM cyclist order by cyclist_name")
	if err != nil {
		fmt.Print(err)
		return []domain.Cyclist{}, err
	}

	for stmt.Next() {
		err = stmt.Scan(&cyclist.Id, &cyclist.IdGroup, &cyclist.Name, &cyclist.Cpf, &cyclist.Birth,
			&cyclist.Email, &cyclist.BloodType, &cyclist.HealthPlan, &cyclist.ContactEmergency, &cyclist.GotToKnow,
			&cyclist.CreateAt, &cyclist.Active, &cyclist.Img, &cyclist.ParticipantType, &cyclist.Pedaling,
			&cyclist.Tours, &cyclist.Travels)
		cyclists = append(cyclists, cyclist)
	}

	return cyclists, nil
}

func (c *CysclistRepositoryPsql) UpdateCyclist(cyclist domain.Cyclist) (domain.Cyclist, error) {

	err := validation.Prepere()
	if err != nil {
		return domain.Cyclist{}, err
	}
	stmt, err := c.db.Prepare("UPDATE cyclist" +
		"SET id_group=$2, cyclist_name=$3, cpf=$4, birth=$5, email=$6, blood_type=$7, health_plan=$8," +
		"contact_emergency=$9, contact_fone=$10, got_to_know=$11, create_at=$12," +
		"active=$13, img=$14, participant_type=$15, pedaling=$16, tours=$17, travels=$18" +
		"WHERE id_cyclist=$1;")
	if err != nil {
		fmt.Print(err)

		return domain.Cyclist{}, err

	}
	_, err = stmt.Exec(cyclist.Id, cyclist.IdGroup, cyclist.Name, cyclist.Cpf, cyclist.Birth, cyclist.Email,
		cyclist.BloodType, cyclist.HealthPlan, cyclist.ContactEmergency, cyclist.GotToKnow,
		cyclist.Active, cyclist.Img, cyclist.ParticipantType, cyclist.Pedaling,
		cyclist.Tours, cyclist.Travels)
	if err != nil {
		return domain.Cyclist{}, err
	}
	err = stmt.Close()
	if err != nil {
		return domain.Cyclist{}, err
	}
	return cyclist, nil
}

func (c *CysclistRepositoryPsql) UpdatePedaling(id string, newPedaling int) (string, error) {
	cli, err := c.Get(id)
	if err != nil {
		fmt.Print(err)
		message := fmt.Sprintf("Hi, %v. deu ruim", err.Error())
		return message, nil
	}
	new_value := cli.Pedaling + newPedaling

	stmt, err := c.db.Prepare("UPDATE cyclist SET  pedaling=$2 WHERE id_cyclist=$1;")
	if err != nil {
		fmt.Print(err)

		return "deu erro", err

	}
	_, err = stmt.Exec(id, new_value)
	if err != nil {
		return "deu ruim", err
	}
	err = stmt.Close()
	if err != nil {
		return "deu ruim", err
	}
	return "pedaldas ataulizadad com sucesso", nil
}

func (c *CysclistRepositoryPsql) UpdateTours(id string, newTour int) (string, error) {
	cli, err := c.Get(id)
	if err != nil {
		fmt.Print(err)
		message := fmt.Sprintf("Hi, %v. deu ruim", err.Error())
		return message, nil
	}
	new_value := cli.Tours + newTour

	stmt, err := c.db.Prepare("UPDATE cyclist SET tours=$2 WHERE id_cyclist=$1;")
	if err != nil {
		fmt.Print(err)

		return "deu erro", err

	}

	_, err = stmt.Exec(id, new_value)
	if err != nil {
		return "deu ruim", err
	}
	err = stmt.Close()
	if err != nil {
		return "deu ruim", err
	}
	return "Passeio ataulizados com sucesso", nil
}

func (c *CysclistRepositoryPsql) UpdateTravels(id string, newTravel int) (string, error) {
	cli, err := c.Get(id)
	if err != nil {
		fmt.Print(err)
		message := fmt.Sprintf("Hi, %v. deu ruim", err.Error())
		return message, nil
	}
	new_value := cli.Travels + newTravel

	stmt, err := c.db.Prepare("UPDATE cyclist SET travels=$2 WHERE id_cyclist=$1;")
	if err != nil {
		fmt.Print(err)

		return "deu erro", err

	}
	_, err = stmt.Exec(id, new_value)
	if err != nil {
		return "deu ruim", err
	}
	err = stmt.Close()
	if err != nil {
		return "deu ruim", err
	}
	return "viagens ataulizadas com sucesso", nil
}

func (g *CysclistRepositoryPsql) Delete(id string) (string, error) {

	stmt, err := g.db.Prepare("DELETE FROM cyclist where id_cyclist=$1")
	if err != nil {
		fmt.Print(err)

		return "erro ao prerar para exclusão", err

	}
	_, err = stmt.Exec(id)
	if err != nil {
		return "Error ao excluir", err
	}
	err = stmt.Close()
	if err != nil {
		return "erro ao fechar conexao", err
	}
	return "Registro excluido com sucesso", nil
}
