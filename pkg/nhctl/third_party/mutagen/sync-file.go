package mutagen

import (
	"fmt"
	"github.com/satori/go.uuid"
	"nocalhost/pkg/nhctl/tools"
)

func FileSync(folder string){
	id := uuid.NewV4()
	idStr := id.String()
	tools.ExecCommand(nil, "mutagen", "sync", "create", "--sync-mode=one-way-safe", fmt.Sprintf("--name=nocalhost-%s", idStr), folder, "root@127.0.0.1:12345:/home/code")
}
