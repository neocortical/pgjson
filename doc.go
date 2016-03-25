/*
Package pgjson is a simple utility package for handling PostgreSQL's JSONB type in Go. Its
code is mostly lifted from code posted by https://github.com/ansel1 in an issue posted to the Gorm
package: https://github.com/jinzhu/gorm/issues/516.

I have added an Unmarshal method as a convenience method for converting into an arbitrary Go struct.
*/

package pgjson
