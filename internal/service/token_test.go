package service

// func TestLogin(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	service := New(repo, nil)
// 	ctx := context.Background()

// 	credentials := "username"
// 	password := "password"
// 	wrongPassword := "wrong_password"

// 	user := new(entity.User)
// 	user.Username = &credentials
// 	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	user.Password.Plaintext = &password
// 	user.Password.Hash = &hash

// 	repo.On("GetUserByCredentials", ctx, credentials).Return(user, nil).Once()
// 	repo.On("CreateToken", ctx, new(util.Token)).Return(nil).Once()
// 	result, err := service.Login(ctx, credentials, password)
// 	assert.NotNil(t, result)
// 	assert.Nil(t, err)

// 	repo.On("GetUserByCredentials", ctx, credentials).Return(nil, errors.New("some error")).Once()
// 	result, err = service.Login(ctx, credentials, password)
// 	assert.Nil(t, result)
// 	assert.NotNil(t, err)

// 	repo.On("GetUserByCredentials", ctx, credentials).Return(user, nil).Once()
// 	result, err = service.Login(ctx, credentials, wrongPassword)
// 	assert.Nil(t, result)
// 	assert.NotNil(t, err)
// }
