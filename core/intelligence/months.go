// SPDX-License-Identifier: AGPL-3.0-or-later

package intelligence

// monthWords lists month spellings that date quotes may use: full names,
// common abbreviations and the inflected forms dates take (Polish, Czech,
// Russian and Ukrainian genitives, Finnish partitives). Index 0 is January.
// Entries are folded like quotes (lowercase, no diacritics) before lookup, and
// a spelling must name the same month in every language that uses it.
// Adding a language means adding one row.
var monthWords = map[string][12][]string{
	"en": {{"january", "jan"}, {"february", "feb"}, {"march", "mar"}, {"april", "apr"}, {"may"}, {"june", "jun"},
		{"july", "jul"}, {"august", "aug"}, {"september", "sep", "sept"}, {"october", "oct"}, {"november", "nov"}, {"december", "dec"}},
	"de": {{"januar", "jänner", "jan"}, {"februar", "feb"}, {"märz", "mär", "mrz"}, {"april", "apr"}, {"mai"}, {"juni", "jun"},
		{"juli", "jul"}, {"august", "aug"}, {"september", "sep", "sept"}, {"oktober", "okt"}, {"november", "nov"}, {"dezember", "dez"}},
	"fr": {{"janvier", "janv"}, {"février", "févr"}, {"mars"}, {"avril", "avr"}, {"mai"}, {"juin"},
		{"juillet", "juil"}, {"août"}, {"septembre", "sept"}, {"octobre", "oct"}, {"novembre", "nov"}, {"décembre", "déc"}},
	"es": {{"enero", "ene"}, {"febrero", "feb"}, {"marzo", "mar"}, {"abril", "abr"}, {"mayo", "may"}, {"junio", "jun"},
		{"julio", "jul"}, {"agosto", "ago"}, {"septiembre", "setiembre", "sep", "sept"}, {"octubre", "oct"}, {"noviembre", "nov"}, {"diciembre", "dic"}},
	"it": {{"gennaio", "gen"}, {"febbraio", "feb"}, {"marzo", "mar"}, {"aprile", "apr"}, {"maggio", "mag"}, {"giugno", "giu"},
		{"luglio", "lug"}, {"agosto", "ago"}, {"settembre", "set"}, {"ottobre", "ott"}, {"novembre", "nov"}, {"dicembre", "dic"}},
	"pt": {{"janeiro", "jan"}, {"fevereiro", "fev"}, {"março", "mar"}, {"abril", "abr"}, {"maio", "mai"}, {"junho", "jun"},
		{"julho", "jul"}, {"agosto", "ago"}, {"setembro", "set"}, {"outubro", "out"}, {"novembro", "nov"}, {"dezembro", "dez"}},
	"nl": {{"januari", "jan"}, {"februari", "feb"}, {"maart", "mrt"}, {"april", "apr"}, {"mei"}, {"juni", "jun"},
		{"juli", "jul"}, {"augustus", "aug"}, {"september", "sep", "sept"}, {"oktober", "okt"}, {"november", "nov"}, {"december", "dec"}},
	"sv": {{"januari"}, {"februari"}, {"mars"}, {"april"}, {"maj"}, {"juni"},
		{"juli"}, {"augusti"}, {"september"}, {"oktober"}, {"november"}, {"december"}},
	"da": {{"januar"}, {"februar"}, {"marts"}, {"april"}, {"maj"}, {"juni"},
		{"juli"}, {"august"}, {"september"}, {"oktober"}, {"november"}, {"december"}},
	"nb": {{"januar"}, {"februar"}, {"mars"}, {"april"}, {"mai"}, {"juni"},
		{"juli"}, {"august"}, {"september"}, {"oktober"}, {"november"}, {"desember", "des"}},
	"fi": {{"tammikuu", "tammikuuta"}, {"helmikuu", "helmikuuta"}, {"maaliskuu", "maaliskuuta"}, {"huhtikuu", "huhtikuuta"},
		{"toukokuu", "toukokuuta"}, {"kesäkuu", "kesäkuuta"}, {"heinäkuu", "heinäkuuta"}, {"elokuu", "elokuuta"},
		{"syyskuu", "syyskuuta"}, {"lokakuu", "lokakuuta"}, {"marraskuu", "marraskuuta"}, {"joulukuu", "joulukuuta"}},
	"pl": {{"styczeń", "stycznia"}, {"luty", "lutego"}, {"marzec", "marca"}, {"kwiecień", "kwietnia"}, {"maj", "maja"},
		{"czerwiec", "czerwca"}, {"lipiec", "lipca"}, {"sierpień", "sierpnia"}, {"wrzesień", "września"},
		{"październik", "października"}, {"listopad", "listopada"}, {"grudzień", "grudnia"}},
	"cs": {{"leden", "ledna"}, {"únor", "února"}, {"březen", "března"}, {"duben", "dubna"}, {"květen", "května"},
		{"červen", "června"}, {"červenec", "července"}, {"srpen", "srpna"}, {"září"}, {"říjen", "října"},
		{"listopad", "listopadu"}, {"prosinec", "prosince"}},
	"tr": {{"ocak"}, {"şubat"}, {"mart"}, {"nisan"}, {"mayıs"}, {"haziran"},
		{"temmuz"}, {"ağustos"}, {"eylül"}, {"ekim"}, {"kasım"}, {"aralık"}},
	"ru": {{"январь", "января"}, {"февраль", "февраля"}, {"март", "марта"}, {"апрель", "апреля"}, {"май", "мая"},
		{"июнь", "июня"}, {"июль", "июля"}, {"август", "августа"}, {"сентябрь", "сентября"}, {"октябрь", "октября"},
		{"ноябрь", "ноября"}, {"декабрь", "декабря"}},
	"uk": {{"січень", "січня"}, {"лютий", "лютого"}, {"березень", "березня"}, {"квітень", "квітня"}, {"травень", "травня"},
		{"червень", "червня"}, {"липень", "липня"}, {"серпень", "серпня"}, {"вересень", "вересня"}, {"жовтень", "жовтня"},
		{"листопад", "листопада"}, {"грудень", "грудня"}},
	"th": {{"มกราคม"}, {"กุมภาพันธ์"}, {"มีนาคม"}, {"เมษายน"}, {"พฤษภาคม"}, {"มิถุนายน"},
		{"กรกฎาคม"}, {"สิงหาคม"}, {"กันยายน"}, {"ตุลาคม"}, {"พฤศจิกายน"}, {"ธันวาคม"}},
}
