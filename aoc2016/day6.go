// --- Day 6: Signals and Noise ---
//
// Something is jamming your communications with Santa. Fortunately, your signal is only partially jammed, and protocol in situations like this is to switch to a simple repetition code to get the message through.
//
// In this model, the same message is sent repeatedly. You've recorded the repeating message signal (your puzzle input), but the data seems quite corrupted - almost too badly to recover. Almost.
//
// All you need to do is figure out which character is most frequent for each position. For example, suppose you had recorded the following messages:
//
// eedadn
// drvtee
// eandsr
// raavrd
// atevrs
// tsrnev
// sdttsa
// rasrtv
// nssdts
// ntnada
// svetve
// tesnvt
// vntsnd
// vrdear
// dvrsen
// enarar
// The most common character in the first column is e; in the second, a; in the third, s, and so on. Combining these characters returns the error-corrected message, easter.
//
// Given the recording in your puzzle input, what is the error-corrected version of the message being sent?
//
// --- Part Two ---
//
// Of course, that would be the message - if you hadn't agreed to use a modified repetition code instead.
//
// In this modified code, the sender instead transmits what looks like random data, but for each character, the character they actually want to send is slightly less likely than the others. Even after signal-jamming noise, you can look at the letter distributions in each column and choose the least common letter to reconstruct the original message.
//
// In the above example, the least common character in the first column is a; in the second, d, and so on. Repeating this process for the remaining characters produces the original message, advent.
//
// Given the recording in your puzzle input and this new decoding methodology, what is the original message that Santa is trying to send?
//
// Your puzzle answer was odqnikqv.

package main

import (
	"fmt"
	"sort"
)

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	input := []string{
		"yksqurlc",
		"hangxyoi",
		"ikyfnszz",
		"lpucpqly",
		"mebitfkg",
		"txvynfxf",
		"qkjjkzlv",
		"zboxjpaz",
		"ijlepotw",
		"rhjolmth",
		"qphnlwok",
		"zhsuzvfc",
		"aqydciuw",
		"xywzachh",
		"jjfprial",
		"fqfigogj",
		"aglymzym",
		"varurjmr",
		"hziejnbj",
		"cuxmzuos",
		"fxlobgee",
		"dmgfztdp",
		"xdycegvz",
		"ovimnzwp",
		"uovtqmsr",
		"ukynvdub",
		"ifsxypxz",
		"bgiindmv",
		"aszamsyt",
		"qhqnxrgw",
		"sikzpkja",
		"adbjyzcm",
		"wegnenxz",
		"yaeqsqlb",
		"xssdsvdd",
		"hfadixdv",
		"xuxilufs",
		"zyulpexe",
		"eyhyrfxg",
		"gzouqych",
		"syswnzvb",
		"hfdfjeky",
		"eueretpq",
		"mawxsflw",
		"abmlttfz",
		"uudtrnrj",
		"hvchepau",
		"uvmmolck",
		"vghofgkt",
		"snxoobnv",
		"teihcctv",
		"owzqcmcu",
		"vagcrkym",
		"awrcvesb",
		"hlrdgrlq",
		"ormblmhn",
		"ottbgibi",
		"liqdgyfl",
		"lqsrrbbm",
		"szvrmevt",
		"yofsqzmw",
		"tsaqhnys",
		"vnvmqzni",
		"sfogzozf",
		"illczjpj",
		"pbaoinrd",
		"tglbyywu",
		"vdzomdgv",
		"wjfwnvhr",
		"knbvicze",
		"efsqjctj",
		"zhpzwpzq",
		"fsstakkb",
		"pgjxrnoo",
		"zppudknc",
		"mwwmlgib",
		"ylufuoqw",
		"wgviqgwg",
		"jsanehvr",
		"jbxvftqj",
		"cgegruav",
		"blvxpgux",
		"rowoutap",
		"zwmlfjks",
		"grptppqn",
		"wqkxnyrd",
		"pufwwpgs",
		"ajccpwgl",
		"kflnghct",
		"dilhwdot",
		"juukkpby",
		"wzvhxrmt",
		"izkfsarl",
		"raschljy",
		"mjjjsied",
		"xqwsdjjn",
		"iftvtwox",
		"buvmkjia",
		"gycuethy",
		"kxgdwhkk",
		"nwvcdkbx",
		"voulrmog",
		"agwgoyce",
		"poaljauy",
		"qaloqwmv",
		"mkhvtlsp",
		"wdfynrws",
		"onqhkohg",
		"slmfuzrb",
		"eihsaceq",
		"pfaxtyha",
		"pdobdtbq",
		"npokorlu",
		"ajlpcoqx",
		"ejxakskg",
		"nvkubumz",
		"bcvfjeya",
		"liozwpey",
		"taukxrgm",
		"mlqefjyg",
		"ctcdrjbm",
		"ujtraucr",
		"ledctflg",
		"ervknztl",
		"ozdtumwv",
		"xxajlviy",
		"duvalefb",
		"plzryrti",
		"ljdhgagu",
		"yfrgqubk",
		"sfbribum",
		"sxehtogn",
		"cckrywmy",
		"jqjtyepp",
		"xvoqwgok",
		"pnfacech",
		"riqzetva",
		"ckqkcggp",
		"auajlcyd",
		"unmamhes",
		"hwgfxfkw",
		"kprsjoar",
		"vuyqxvcg",
		"vtnqfvyw",
		"aptjqhji",
		"vxurgkpn",
		"jcimbhpd",
		"flowelzo",
		"egkvzpnn",
		"gpdyzsue",
		"liecisqv",
		"uzzujvre",
		"vnwetclx",
		"sebvvzta",
		"bcysxtaq",
		"gskvmoec",
		"dlwjzfsv",
		"jzeuaqjo",
		"wqfxxqhl",
		"debxresp",
		"nlmvrgnl",
		"qlagubug",
		"hgzydwon",
		"riyyoamh",
		"dttgbion",
		"fkxdjaau",
		"jacyfdqr",
		"nqblwnjj",
		"hjqkhqhx",
		"rsdedeww",
		"igjlswcl",
		"efrsumkl",
		"yqktlclb",
		"kwktcqfm",
		"yxfxrukc",
		"aymhfxnw",
		"wwucvnks",
		"evedwqxu",
		"yyxuioac",
		"jrkpkxfm",
		"lmvzbbcv",
		"ccnchsxa",
		"hrczhzsu",
		"ywvtltuh",
		"tbuuaxwm",
		"mojysgwi",
		"edrndvyd",
		"khevyajf",
		"ebipchau",
		"urxhjxpf",
		"idbgidtp",
		"kqizxqie",
		"rrdwymmp",
		"exzcgcdd",
		"yypeeexx",
		"yhbxpbgj",
		"xtxwxtqh",
		"zvkjhjsj",
		"kvdzbped",
		"txrgskdd",
		"jlyiktsr",
		"ukiefbpz",
		"zlovjsla",
		"odnfxujt",
		"qaeeskke",
		"lveabhde",
		"cjafenie",
		"xqjsxxhg",
		"fpdkojtp",
		"picajdgr",
		"thzbrksz",
		"jgalucac",
		"xuznhoqk",
		"sdvndifi",
		"rshebeoe",
		"bmawstny",
		"cdabhicz",
		"pegfhana",
		"talouldl",
		"rrkrqnic",
		"jvuyppxo",
		"qmpvtnho",
		"uivhaphi",
		"jtojvmbm",
		"dmgukifi",
		"yfqaakrp",
		"rdgbqoep",
		"brmwsyol",
		"esymvcsm",
		"ibzjyrke",
		"oardzxmx",
		"mlurkbwc",
		"dpiecmuf",
		"qxhayxje",
		"lluwcpug",
		"iqaryifh",
		"zqsvpabf",
		"axctuzab",
		"srmjhzzl",
		"kmekssik",
		"pdlkpwfz",
		"ggfeijyk",
		"briuewuu",
		"tyymvfkl",
		"nnrhosbw",
		"uftcqtpu",
		"xpasvqvx",
		"fkhpevzo",
		"tvbphorw",
		"jolqmzyj",
		"hempphpk",
		"xtphbkpi",
		"kuroidgo",
		"cmxzeejx",
		"lutcewnx",
		"obqqiyoj",
		"difqxijk",
		"qissebkj",
		"tuplgcyq",
		"ksftyayf",
		"efxbefvq",
		"deojflsk",
		"chsdnamb",
		"ingogbdt",
		"cigvurzg",
		"cfnkgoro",
		"grcqjxbk",
		"xbbofyif",
		"mbpyvzqo",
		"naugflph",
		"spbcqwtf",
		"ulbrfapx",
		"gwjwlvif",
		"bdrhzxnd",
		"rcjwdapc",
		"dktbilqi",
		"qmnfwejv",
		"twqqfdza",
		"wkwoabmb",
		"fqcpaxys",
		"iindosuu",
		"pbafpsfg",
		"zbkvbhgo",
		"mmynvdoq",
		"rffzwlsv",
		"nyuwkijl",
		"afnxwxst",
		"zmbufurn",
		"yggxflsq",
		"narogdzl",
		"qzsqcmrn",
		"oaghbhvp",
		"vneroflt",
		"qklqllzh",
		"ngwppigb",
		"ldqwjrkf",
		"fkyjlvan",
		"vocndsst",
		"cmyxgnks",
		"myqmairp",
		"szojbcwf",
		"pzqytons",
		"gmekbavj",
		"tijaeqxb",
		"iyntdjqj",
		"osnsslss",
		"svxqpvvs",
		"rcoyrlvq",
		"scznwrsz",
		"mfccxcqm",
		"ujsuzcxu",
		"fqkvvgab",
		"pvemyiim",
		"hljfqgvn",
		"jodbyaod",
		"ijyydnhy",
		"hgejzxvs",
		"widmvhho",
		"lupizvbi",
		"aaqgqzjq",
		"bblprnpa",
		"tezfndcz",
		"qkmamabn",
		"rxnwxeya",
		"clvourij",
		"efhxuazl",
		"ctugmmik",
		"pdbshihk",
		"siaubyzi",
		"fmkxcsdt",
		"qhsomhyx",
		"rxlbiigd",
		"wksiikcr",
		"utcmzwoo",
		"fcxcvtdj",
		"tfdnnjxi",
		"zpjusqzn",
		"xexktjpi",
		"nizpdylh",
		"xqphsati",
		"qgzrgzwc",
		"pddtewuz",
		"vblmyybo",
		"zjzraydp",
		"grhvkkff",
		"ixhmhtly",
		"oowiznms",
		"vfjpnlju",
		"gpjkuzie",
		"uisyhxdm",
		"gcjlougf",
		"abyunjpn",
		"hndfklet",
		"cznialyv",
		"xzdqdsqh",
		"ideijfqq",
		"wlsekkwm",
		"mmnkpbxu",
		"eopdnmhd",
		"hhtfqgvc",
		"ihluufyc",
		"idclmjin",
		"jilmkrnx",
		"fudjusjh",
		"lefmnbqc",
		"vcoohrwk",
		"qnbwqvgr",
		"ajhavirh",
		"ztwszgds",
		"thhmfcuf",
		"yxuxodnk",
		"lffzrzlk",
		"wnpiyutz",
		"bgqbitvg",
		"juhdxwhr",
		"zcsvtbtq",
		"kxbmtebv",
		"guzhbyct",
		"yvghwoco",
		"ilqnmxjt",
		"wwynowzy",
		"hettyliu",
		"gknfnjgb",
		"fmqvbnws",
		"dxcmwibd",
		"eqindscs",
		"qntvuhor",
		"bmiytxhp",
		"szvbgkjb",
		"oohlbhij",
		"cajumomc",
		"xglxwpxl",
		"icckprut",
		"fexesnsy",
		"ddjzlmex",
		"yjkiajrx",
		"hiwbtwxr",
		"loutjfzj",
		"lyxuomov",
		"lrvgcyxp",
		"kjxgyrpn",
		"xookinff",
		"hskiiitj",
		"stsfcxfi",
		"uxfuafbb",
		"umagoczd",
		"hwepkhqd",
		"ksljzsxu",
		"lzzdnxvi",
		"xbnwogpd",
		"wjljquis",
		"mcuezdet",
		"rjjlgwge",
		"bgpexmex",
		"ktugiefo",
		"gspmsutg",
		"zmooklqs",
		"essgayfk",
		"ewmvttkc",
		"nlxqluzn",
		"vzbzpssa",
		"nttyoaxz",
		"ghymysoh",
		"pfrpihxw",
		"womaopjr",
		"jcohbgwl",
		"gcxtvhze",
		"crdckley",
		"ybawrahw",
		"qyzjvttl",
		"njcgqexz",
		"wuhhymlq",
		"qamapddz",
		"obritukf",
		"nrkzgzuh",
		"anrxndlr",
		"bkhhwpns",
		"anwgnwxs",
		"oryzjgrw",
		"vyisulmw",
		"owqtcrrp",
		"ztwbxmjk",
		"uncbpjdm",
		"kwpjlomg",
		"bhjllqlg",
		"vyetfpwo",
		"gpkvdoym",
		"mhvqraey",
		"ibuauftc",
		"dszirmit",
		"rzprzetb",
		"dutethfs",
		"rncgofwb",
		"ttjpagvt",
		"zhgpfuun",
		"jrmbbcfa",
		"darbwwvw",
		"gpfiaicg",
		"nzgsddrf",
		"zyxlhgvm",
		"fxftxkuq",
		"wmbtdlra",
		"bnnwksww",
		"lnadfibo",
		"bwtvzqax",
		"fchlbbhy",
		"etgvnukq",
		"phwqlwyr",
		"usmqwyuy",
		"phiywjtf",
		"vquksvbc",
		"herdwxwy",
		"tvrelfhx",
		"mkystkee",
		"wlkzgemz",
		"tvwzmzbk",
		"dcizxyom",
		"occbsbrc",
		"bgtejfka",
		"pezzjvww",
		"dobsvxeq",
		"suxsxupg",
		"awtdvqqk",
		"aadicudh",
		"osmszjdi",
		"dwiruegd",
		"gaptsmnw",
		"xptoqiiy",
		"czrwwtcj",
		"fuaojfef",
		"legftmnv",
		"utirkhom",
		"wqztlrav",
		"nqrsdnel",
		"fphioxvo",
		"ynpqmvqq",
		"fgpkmqqd",
		"syqkahmr",
		"nicgaplr",
		"vboaxnid",
		"kvtuiumu",
		"yygccoaa",
		"qorndtpx",
		"vkmnvino",
		"nhtyablj",
		"svheacne",
		"ccggmhlg",
		"dwjbuwli",
		"jelrgful",
		"valsexan",
		"rlpydgeq",
		"crwdruho",
		"umwacydj",
		"zmkhupee",
		"fduhhheq",
		"oebabkhi",
		"fbzeojzp",
		"tovjquoa",
		"kqvlbsnz",
		"eorjsjau",
		"xcmcrvxh",
		"mttahrcg",
		"pptpevct",
		"veeacqsu",
		"hbicddre",
		"mwmlikcb",
		"kdkwgbtf",
		"gzbfcoqm",
		"ryakvpdc",
		"njxwxgjm",
		"jpfkjwwa",
		"ucilydrh",
		"avfqktrr",
		"czsxmunf",
		"mrtigjvl",
		"msqdtrao",
		"hnxnhgaq",
		"bzywgvsv",
		"qvinjvew",
		"wnvpvswq",
		"swqbbypt",
		"qgnlkafy",
		"whdxfmba",
		"ijnuvztp",
		"dpgpstgz",
		"mphsexii",
		"zhdbcqwz",
		"ereaqdsh",
		"fdbllkue",
		"rxnhmffr",
		"dtslnbqv",
		"rapwovhx",
		"kseekmuv",
		"yvuxmqon",
		"myjxrgmu",
		"bkdocnae",
		"wwfrldmd",
		"tsccmqnn",
		"pemyiddw",
		"oxnaefeh",
		"dtodqyao",
		"loglzqcp",
		"beidpoza",
		"nhwppnyb",
		"izwzhrnk",
		"seefzfdf",
		"axtawcjg",
		"htozfqfr",
		"buoiqrbu",
		"gvpsnstc",
		"lkayybuc",
		"kywnhosu",
		"gsdzlevi",
		"nvfefkpa",
		"koineqzp",
		"njqopqrt",
		"jkhpcbkh",
		"ypighaih",
		"ytgiblyn",
		"bscfyvil",
		"jknjtpmr",
		"qqoroalx",
		"kmymocvo",
		"xookuydy",
		"oyncmltd",
		"pmgsgpgg",
		"zhyrmznp",
		"tcfysexy",
		"momevddz",
		"xrkiiwzb",
		"uiybjhpx",
		"cxzphbgw",
		"eqvrtbkt",
		"rtetdnyc",
		"zbwpwkve",
		"arhokpma",
		"yspdacfj",
		"opnkfcjk"}

	length := len(input)
	var str string
	for i := 0; i < 8; i++ {
		count := make(map[string]int)
		for j := 0; j < length; j++ {
			byt := []byte(input[j])
			count[string(byt[i])] = count[string(byt[i])] + 1
		}
		pl := rankByWordCount(count)
		str = str + pl[0].Key
	}

	fmt.Println(str)
}

// part two: in function Less, change from "<" to ">"
