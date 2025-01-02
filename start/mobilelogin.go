package start

//"regexp"

//const otpLength = 4

// Function to generate a random OTP of given length
/*func generateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	min := int64(pow10(length - 1))
	max := int64(pow10(length) - 1)
	otp := rand.Int63n(max-min+1) + min
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"d", otp)
}*/

/* func generateOTP(length int) int {
	rand.Seed(time.Now().UnixNano())
	min := int(pow10(length - 1))
	max := int(pow10(length) - 1)
	otp := rand.Intn(max-min+1) + min
	return otp
}

// Helper function to calculate the power of 10
func pow10(n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
*/
// Function to validate if the mobile number is 10 digits
/* func validateMobileNumber(mobileNumber int64) bool {
	mobileNumberPattern := regexp.MustCompile(`^\d{10}$`)
	return mobileNumberPattern.MatchString(strconv.FormatInt(mobileNumber, 10))
}
*/
/* func convertToInt32(num int) int32 {
	return int32(num)
}
*/
// Function to create a mobile OTP and update the database based on conditions
/* func CreateMobileOTP(client *ent.Client, newadminlogin *ent.AdminLogin) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Retrieve the employee profile by ID and username
	employeeProfile, err := client.AdminLogin.
		Query().
		Where(adminlogin.Username(newadminlogin.Username)).
		Only(ctx)
	if err != nil {
		return "", err
	}

	// Check if the employee profile exists
	if employeeProfile == nil {
		return "", errors.New(" employee profile does not exist")
	}

	// Check if the role matches
	if employeeProfile.RoleName != newadminlogin.RoleName {
		return "", errors.New(" role does not match")
	}

	// Check for mobile Number
	if employeeProfile.MobileNumber == 0 {
		return "", errors.New(" no mobile number in employee profile")
	} else if employeeProfile.MobileNumber > 0 {

		// Convert the mobile number to string and check if it is of 10 digits
		mobileNumberStr := strconv.FormatInt(employeeProfile.MobileNumber, 10)
		if len(mobileNumberStr) != 10 {
			return "", errors.New("  mobile number is not of 10 digits")
		}

		// Generate OTP
		otp := generateOTP(4)
		otpInt32 := convertToInt32(otp)
		//otp = convertToInt32(otp)

		// Update the OTP and verify remarks in the AdminLogin entity

		_, err = employeeProfile.Update().
			SetOTP(otpInt32).
			SetVerifyRemarks("OTP generated").
			Save(ctx)

		if err != nil {
			return "", err
		}

		// Send OTP to the mobile number
		//err = sendOTP(employeeProfile.MobileNumber, otp)
		//if err != nil {
		//	return "", err
		//}

		// Print the username and generated OTP
		fmt.Printf("Username: %s\n", newadminlogin.Username)
		//fmt.Printf("Generated OTP for EMP ID %d is: %s\n", newadminlogin.Username, otp)
		fmt.Printf("Generated OTP for EMP ID %s is: %d\n", newadminlogin.Username, otp)

		// Return the generated OTP
		return "happy", nil
	}

	return "", errors.New(" invalid mobile number")
} */
