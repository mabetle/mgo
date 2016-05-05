// sdb define some APIs
//
// SimpleTable
// a csv(or tsv), sheet, database table can abstract to SimpleTable
// the ideas from jdbc API
// Next() for loop, and holder a rowIndex
//
// RandomAccessTable
// There is no rowIndex concept, but you can walk rows too with loop.
// its very convinient for access sheet when you know its struct.
// for example: import values from a fix format sheet.
package msdb
