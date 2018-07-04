// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package two

type Pagination struct {
	Page  int `url:"page"`
	Limit int `url:"limit"`
}
