// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package shadow_drive_user_staking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateAccount is the `updateAccount` instruction.
type UpdateAccount struct {
	Identifier *string                `bin:"optional"`
	Owner2     *ag_solanago.PublicKey `bin:"optional"`

	// [0] = [] storageConfig
	//
	// [1] = [WRITE] storageAccount
	//
	// [2] = [WRITE, SIGNER] owner
	//
	// [3] = [] tokenMint
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateAccountInstructionBuilder creates a new `UpdateAccount` instruction builder.
func NewUpdateAccountInstructionBuilder() *UpdateAccount {
	nd := &UpdateAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetIdentifier sets the "identifier" parameter.
func (inst *UpdateAccount) SetIdentifier(identifier string) *UpdateAccount {
	inst.Identifier = &identifier
	return inst
}

// SetOwner2 sets the "owner2" parameter.
func (inst *UpdateAccount) SetOwner2(owner2 ag_solanago.PublicKey) *UpdateAccount {
	inst.Owner2 = &owner2
	return inst
}

// SetStorageConfigAccount sets the "storageConfig" account.
func (inst *UpdateAccount) SetStorageConfigAccount(storageConfig ag_solanago.PublicKey) *UpdateAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(storageConfig)
	return inst
}

// GetStorageConfigAccount gets the "storageConfig" account.
func (inst *UpdateAccount) GetStorageConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStorageAccountAccount sets the "storageAccount" account.
func (inst *UpdateAccount) SetStorageAccountAccount(storageAccount ag_solanago.PublicKey) *UpdateAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(storageAccount).WRITE()
	return inst
}

// GetStorageAccountAccount gets the "storageAccount" account.
func (inst *UpdateAccount) GetStorageAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *UpdateAccount) SetOwnerAccount(owner ag_solanago.PublicKey) *UpdateAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *UpdateAccount) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *UpdateAccount) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *UpdateAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *UpdateAccount) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UpdateAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UpdateAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.StorageConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.StorageAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UpdateAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Identifier (OPT)", inst.Identifier))
						paramsBranch.Child(ag_format.Param("    Owner2 (OPT)", inst.Owner2))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("storageConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      storage", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    tokenMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UpdateAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Identifier` param (optional):
	{
		if obj.Identifier == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Identifier)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Owner2` param (optional):
	{
		if obj.Owner2 == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Owner2)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *UpdateAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Identifier` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Identifier)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Owner2` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Owner2)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewUpdateAccountInstruction declares a new UpdateAccount instruction with the provided parameters and accounts.
func NewUpdateAccountInstruction(
	// Parameters:
	identifier string,
	owner2 ag_solanago.PublicKey,
	// Accounts:
	storageConfig ag_solanago.PublicKey,
	storageAccount ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UpdateAccount {
	return NewUpdateAccountInstructionBuilder().
		SetIdentifier(identifier).
		SetOwner2(owner2).
		SetStorageConfigAccount(storageConfig).
		SetStorageAccountAccount(storageAccount).
		SetOwnerAccount(owner).
		SetTokenMintAccount(tokenMint).
		SetSystemProgramAccount(systemProgram)
}
