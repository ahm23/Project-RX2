// x/storage/keeper/file.go
package keeper

import (
	"context"

	"nebulix/x/storage/types"
)

// StoreFile stores a new file with a provider
func (k Keeper) SetFile(ctx context.Context, file types.File) error {
	// sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if file already exists
	exists, err := k.FileStore.Has(ctx, file.Merkle)
	if err != nil {
		return err
	}
	if exists {
		return types.ErrFileAlreadyExists
	}

	// Check if provider exists and is active
	provider, err := k.ProviderStore.Get(ctx, file.Providers[0])
	if err != nil {
		return types.ErrProviderNotFound
	}
	// if provider.Status != types.ProviderStatusActive {
	// 	return types.ErrProviderNotActive
	// }

	// MOVE TO MSG SERVER
	// // Create file
	// file := types.File{
	// 	Hash:       fileId,
	// 	FileSize:   size,
	// 	Providers:  []string{providerAddr},
	// 	StartBlock: sdkCtx.BlockHeight(),
	// }

	// Save file
	if err := k.FileStore.Set(ctx, file.Merkle, file); err != nil {
		return err
	}

	// TODO: update provider space used
	// provider.TotalFiles++
	if err := k.ProviderStore.Set(ctx, provider.Address, provider); err != nil {
		return err
	}

	// Add to provider index
	// providerFileKey := collections.Join(providerAddr, fileId)
	// if err := k.FilesByProvider.Set(ctx, providerFileKey); err != nil {
	// 	return err
	// }

	return nil
}

// \\ GetFile retrieves a file by ID
func (k Keeper) GetFile(ctx context.Context, fileId string) (types.File, error) {
	return k.FileStore.Get(ctx, fileId)
}

// \\ DeleteFile marks a file as deleted
func (k Keeper) DeleteFile(ctx context.Context, fileId string, owner string) error {
	file, err := k.FileStore.Get(ctx, fileId)
	if err != nil {
		return err
	}

	// TODO: Verify ownership?

	// Soft delete
	// file.IsDeleted = true
	return k.FileStore.Set(ctx, fileId, file)
}

func (k Keeper) IterateAllFiles(ctx context.Context, fn func(file types.File) (stop bool)) error {
	iter, err := k.FileStore.Iterate(ctx, nil)
	if err != nil {
		return err
	}
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		kv, err := iter.KeyValue()
		if err != nil {
			// [TBD]: continue?
			return err
		}

		if fn(kv.Value) {
			break
		}
	}

	return nil
}
