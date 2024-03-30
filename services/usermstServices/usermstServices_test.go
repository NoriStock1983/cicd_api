package usermstservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUserList(t *testing.T) {

	expected := "User List!"

	/* actualに、usermstService.go内のGetUserList関数の戻り値をセットする。 */
	actual := GetUserList()

	assert.Equal(t, expected, actual)

}
