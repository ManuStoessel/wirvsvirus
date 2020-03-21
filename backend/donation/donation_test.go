package donation

import (
	"testing"

	"github.com/ManuStoessel/wirvsvirus/backend/user"
	User "github.com/ManuStoessel/wirvsvirus/backend/user"
)

var testEmail string = "info@alexanderwagner.eu"
var testName string = "Alex"

func Test(t *testing.T) {
	user := &user.User{Username: testName, Email: testEmail}
	User.Create(user)

	userID := user.ID

	donation := &Donation{ReceiverID: userID, Amount: 10.00}
	donation2 := &Donation{ReceiverID: userID, Amount: 20.00}

	Create(donation)
	Create(donation2)

	donationID := donation.ID

	readDonation := Read(donationID)

	if readDonation.ID != donation.ID {
		t.Errorf("Read donation should be equal created donation")
	}

	donation.Amount = 15.00

	Update(donation)

	readDonation = Read(donationID)

	if readDonation.Amount != 15.00 {
		t.Errorf("Amount should be changed")
	}

	donations := ReadByUser(userID)

	if len(donations) != 2 {
		t.Errorf("There should be 2 donations")
	}

	Delete(donation)

	readDonation = Read(donationID)

	if readDonation.ID != "" {
		t.Errorf("Donation should not exist anymore")
	}
}
