// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package shadow_drive_user_staking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UnmarkDeleteAccount2 is the `unmarkDeleteAccount2` instruction.
type UnmarkDeleteAccount2 struct {

	// [0] = [] storageConfig
	//
	// [1] = [WRITE] storageAccount
	//
	// [2] = [WRITE] stakeAccount
	//
	// [3] = [WRITE, SIGNER] owner
	//
	// [4] = [] tokenMint
	//
	// [5] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUnmarkDeleteAccount2InstructionBuilder creates a new `UnmarkDeleteAccount2` instruction builder.
func NewUnmarkDeleteAccount2InstructionBuilder() *UnmarkDeleteAccount2 {
	nd := &UnmarkDeleteAccount2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetStorageConfigAccount sets the "storageConfig" account.
func (inst *UnmarkDeleteAccount2) SetStorageConfigAccount(storageConfig ag_solanago.PublicKey) *UnmarkDeleteAccount2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(storageConfig)
	return inst
}

// GetStorageConfigAccount gets the "storageConfig" account.
func (inst *UnmarkDeleteAccount2) GetStorageConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStorageAccountAccount sets the "storageAccount" account.
func (inst *UnmarkDeleteAccount2) SetStorageAccountAccount(storageAccount ag_solanago.PublicKey) *UnmarkDeleteAccount2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(storageAccount).WRITE()
	return inst
}

// GetStorageAccountAccount gets the "storageAccount" account.
func (inst *UnmarkDeleteAccount2) GetStorageAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStakeAccountAccount sets the "stakeAccount" account.
func (inst *UnmarkDeleteAccount2) SetStakeAccountAccount(stakeAccount ag_solanago.PublicKey) *UnmarkDeleteAccount2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccountAccount gets the "stakeAccount" account.
func (inst *UnmarkDeleteAccount2) GetStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOwnerAccount sets the "owner" account.
func (inst *UnmarkDeleteAccount2) SetOwnerAccount(owner ag_solanago.PublicKey) *UnmarkDeleteAccount2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *UnmarkDeleteAccount2) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *UnmarkDeleteAccount2) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *UnmarkDeleteAccount2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *UnmarkDeleteAccount2) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UnmarkDeleteAccount2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UnmarkDeleteAccount2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UnmarkDeleteAccount2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst UnmarkDeleteAccount2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UnmarkDeleteAccount2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UnmarkDeleteAccount2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UnmarkDeleteAccount2) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.StorageConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.StorageAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.StakeAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UnmarkDeleteAccount2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UnmarkDeleteAccount2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("storageConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      storage", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        stake", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    tokenMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj UnmarkDeleteAccount2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UnmarkDeleteAccount2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUnmarkDeleteAccount2Instruction declares a new UnmarkDeleteAccount2 instruction with the provided parameters and accounts.
func NewUnmarkDeleteAccount2Instruction(
	// Accounts:
	storageConfig ag_solanago.PublicKey,
	storageAccount ag_solanago.PublicKey,
	stakeAccount ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UnmarkDeleteAccount2 {
	return NewUnmarkDeleteAccount2InstructionBuilder().
		SetStorageConfigAccount(storageConfig).
		SetStorageAccountAccount(storageAccount).
		SetStakeAccountAccount(stakeAccount).
		SetOwnerAccount(owner).
		SetTokenMintAccount(tokenMint).
		SetSystemProgramAccount(systemProgram)
}
