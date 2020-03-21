package entity

import (
	"testing"
)

func TestDonation(t *testing.T) {
	var testEmail string = "info@alexanderwagner.eu"
	var testName string = "Alex"

	user := &User{Username: testName, Email: testEmail}
	user.Create()

	userID := user.ID

	donation := &Donation{ReceiverID: userID, Amount: 10.00}
	donation2 := &Donation{ReceiverID: userID, Amount: 20.00}

	donation.Create()
	donation2.Create()

	donationID := donation.ID

	readDonation := &Donation{}
	readDonation = readDonation.Read(donationID)

	if readDonation.ID != donation.ID {
		t.Errorf("Read donation should be equal created donation")
	}

	donation.Amount = 15.00

	donation.Update()

	readDonation = readDonation.Read(donationID)

	if readDonation.Amount != 15.00 {
		t.Errorf("Amount should be changed")
	}

	donations := readDonation.ReadByUser(userID)

	if len(donations) != 2 {
		t.Errorf("There should be 2 donations")
	}

	donation.Delete()

	readDonation = readDonation.Read(donationID)

	if readDonation.ID != "" {
		t.Errorf("Donation should not exist anymore")
	}
}
