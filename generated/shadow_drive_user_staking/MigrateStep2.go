// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package shadow_drive_user_staking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MigrateStep2 is the `migrateStep2` instruction.
type MigrateStep2 struct {

	// [0] = [WRITE] storageAccount
	//
	// [1] = [WRITE] migration
	//
	// [2] = [WRITE, SIGNER] owner
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMigrateStep2InstructionBuilder creates a new `MigrateStep2` instruction builder.
func NewMigrateStep2InstructionBuilder() *MigrateStep2 {
	nd := &MigrateStep2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetStorageAccountAccount sets the "storageAccount" account.
func (inst *MigrateStep2) SetStorageAccountAccount(storageAccount ag_solanago.PublicKey) *MigrateStep2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(storageAccount).WRITE()
	return inst
}

// GetStorageAccountAccount gets the "storageAccount" account.
func (inst *MigrateStep2) GetStorageAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMigrationAccount sets the "migration" account.
func (inst *MigrateStep2) SetMigrationAccount(migration ag_solanago.PublicKey) *MigrateStep2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(migration).WRITE()
	return inst
}

// GetMigrationAccount gets the "migration" account.
func (inst *MigrateStep2) GetMigrationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *MigrateStep2) SetOwnerAccount(owner ag_solanago.PublicKey) *MigrateStep2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *MigrateStep2) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MigrateStep2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MigrateStep2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MigrateStep2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst MigrateStep2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MigrateStep2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MigrateStep2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MigrateStep2) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.StorageAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Migration is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *MigrateStep2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MigrateStep2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      storage", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    migration", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj MigrateStep2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *MigrateStep2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewMigrateStep2Instruction declares a new MigrateStep2 instruction with the provided parameters and accounts.
func NewMigrateStep2Instruction(
	// Accounts:
	storageAccount ag_solanago.PublicKey,
	migration ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *MigrateStep2 {
	return NewMigrateStep2InstructionBuilder().
		SetStorageAccountAccount(storageAccount).
		SetMigrationAccount(migration).
		SetOwnerAccount(owner).
		SetSystemProgramAccount(systemProgram)
}
