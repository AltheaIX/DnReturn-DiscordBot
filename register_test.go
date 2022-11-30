package main

/* func TestDoRegister(t *testing.T) {
	payload := "suis!prefix HaelaCi:Haela123:777777:HaelaChiii@gmail.com"
	regex, _ := regexp.Compile(`[^\s]*$`)
	regexPayload := regex.FindAllString(payload, -1)
	splitPayload := strings.Split(regexPayload[0], ":")
	username, password, pin, email := splitPayload[0], splitPayload[1], splitPayload[2], splitPayload[3]
	register := doRegister(username, password, pin, email)
	fmt.Println(register)
	fmt.Println(splitPayload)
}

func TestNameGenApi(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fakeName := nameGenApi()
	caser := cases.Title(language.English)
	username := caser.String(fakeName.Username)
	password := fmt.Sprintf("%v%v", caser.String(fakeName.Password), (rand.Intn(9999))+100)
	fmt.Printf("username: %v, password: %v", username, password)
}

func TestDoRandomRegister(t *testing.T) {
	result, payload := doRandomRegister(2)
	fmt.Println(result)
	fmt.Println(payload)
} */
