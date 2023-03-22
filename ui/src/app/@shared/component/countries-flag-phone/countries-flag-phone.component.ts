import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'country-phone',
  templateUrl: './countries-flag-phone.component.html',
  styleUrls: ['./countries-flag-phone.component.scss'],
})
export class CountriesFlagPhone implements OnInit {
  countries = {
    'af': {'code': 'af', 'cn': '阿富汗', 'en': 'Afghanistan (‫افغانستان‬‎)', 'dialCode': 93, 'phoneFormat': '070 123 4567'},
    'al': {'code': 'al', 'cn': '阿尔巴尼亚', 'en': 'Albania (Shqipëri)', 'dialCode': 355, 'phoneFormat': '066 123 4567'},
    'dz': {'code': 'dz', 'cn': '阿尔及利亚', 'en': 'Algeria (‫الجزائر‬‎)', 'dialCode': 213, 'phoneFormat': '0551 23 45 67'},
    'as': {'code': 'as', 'cn': '美属萨摩亚', 'en': 'American Samoa', 'dialCode': 1684, 'phoneFormat': '(684) 733-1234'},
    'ad': {'code': 'ad', 'cn': '安道尔', 'en': 'Andorra', 'dialCode': 376, 'phoneFormat': '312 345'},
    'ao': {'code': 'ao', 'cn': '安哥拉', 'en': 'Angola', 'dialCode': 244, 'phoneFormat': '923 123 456'},
    'ai': {'code': 'ai', 'cn': '安圭拉', 'en': 'Anguilla', 'dialCode': 1264, 'phoneFormat': '(264) 235-1234'},
    'ag': {'code': 'ag', 'cn': '安提瓜和巴布达', 'en': 'Antigua and Barbuda', 'dialCode': 1268, 'phoneFormat': '(268) 464-1234'},
    'ar': {'code': 'ar', 'cn': '阿根廷', 'en': 'Argentina', 'dialCode': 54, 'phoneFormat': '011 15-2345-6789'},
    'am': {'code': 'am', 'cn': '亚美尼亚', 'en': 'Armenia (Հայաստան)', 'dialCode': 374, 'phoneFormat': '077 123456'},
    'aw': {'code': 'aw', 'cn': '阿鲁巴', 'en': 'Aruba', 'dialCode': 297, 'phoneFormat': '560 1234'},
    'au': {'code': 'au', 'cn': '澳大利亚', 'en': 'Australia', 'dialCode': 61, 'phoneFormat': '0412 345 678'},
    'at': {'code': 'at', 'cn': '奥地利', 'en': 'Austria (Österreich)', 'dialCode': 43, 'phoneFormat': '0664 123456'},
    'az': {'code': 'az', 'cn': '阿塞拜疆', 'en': 'Azerbaijan (Azərbaycan)', 'dialCode': 994, 'phoneFormat': '040 123 45 67'},
    'bs': {'code': 'bs', 'cn': '巴哈马', 'en': 'Bahamas', 'dialCode': 1242, 'phoneFormat': '(242) 359-1234'},
    'bh': {'code': 'bh', 'cn': '巴林', 'en': 'Bahrain (‫البحرين‬‎)', 'dialCode': 973, 'phoneFormat': '3600 1234'},
    'bd': {'code': 'bd', 'cn': '孟加拉', 'en': 'Bangladesh (বাংলাদেশ)', 'dialCode': 880, 'phoneFormat': '01812-345678'},
    'bb': {'code': 'bb', 'cn': '巴巴多斯', 'en': 'Barbados', 'dialCode': 1246, 'phoneFormat': '(246) 250-1234'},
    'by': {'code': 'by', 'cn': '白俄罗斯共和国', 'en': 'Belarus (Беларусь)', 'dialCode': 375, 'phoneFormat': '8 029 491-19-11'},
    'be': {'code': 'be', 'cn': '比利时', 'en': 'Belgium (België)', 'dialCode': 32, 'phoneFormat': '0470 12 34 56'},
    'bz': {'code': 'bz', 'cn': '伯利兹', 'en': 'Belize', 'dialCode': 501, 'phoneFormat': '622-1234'},
    'bj': {'code': 'bj', 'cn': '贝宁', 'en': 'Benin (Bénin)', 'dialCode': 229, 'phoneFormat': '90 01 12 34'},
    'bm': {'code': 'bm', 'cn': '百慕大', 'en': 'Bermuda', 'dialCode': 1441, 'phoneFormat': '(441) 370-1234'},
    'bt': {'code': 'bt', 'cn': '不丹', 'en': 'Bhutan (འབྲུག)', 'dialCode': 975, 'phoneFormat': '17 12 34 56'},
    'bo': {'code': 'bo', 'cn': '玻利维亚', 'en': 'Bolivia', 'dialCode': 591, 'phoneFormat': '71234567'},
    'ba': {
      'code': 'ba',
      'cn': '波黑',
      'en': 'Bosnia and Herzegovina (Босна и Херцеговина)',
      'dialCode': 387,
      'phoneFormat': '061 123 456'
    },
    'bw': {'code': 'bw', 'cn': '博茨瓦纳', 'en': 'Botswana', 'dialCode': 267, 'phoneFormat': '71 123 456'},
    'br': {'code': 'br', 'cn': '伯利兹', 'en': 'Brazil (Brasil)', 'dialCode': 55, 'phoneFormat': '(11) 96123-4567'},
    'io': {'code': 'io', 'cn': '英属印度洋领地', 'en': 'British Indian Ocean Territory', 'dialCode': 246, 'phoneFormat': '380 1234'},
    'vg': {'code': 'vg', 'cn': '英属维尔京群岛', 'en': 'British Virgin Islands', 'dialCode': 1284, 'phoneFormat': '(284) 300-1234'},
    'bn': {'code': 'bn', 'cn': '文莱', 'en': 'Brunei', 'dialCode': 673, 'phoneFormat': '712 3456'},
    'bg': {'code': 'bg', 'cn': '保加利亚', 'en': 'Bulgaria (България)', 'dialCode': 359, 'phoneFormat': '048 123 456'},
    'bf': {'code': 'bf', 'cn': '布基纳法索', 'en': 'Burkina Faso', 'dialCode': 226, 'phoneFormat': '70 12 34 56'},
    'bi': {'code': 'bi', 'cn': '布隆迪', 'en': 'Burundi (Uburundi)', 'dialCode': 257, 'phoneFormat': '79 56 12 34'},
    'kh': {'code': 'kh', 'cn': '柬埔寨', 'en': 'Cambodia (កម្ពុជា)', 'dialCode': 855, 'phoneFormat': '091 234 567'},
    'cm': {'code': 'cm', 'cn': '喀麦隆', 'en': 'Cameroon (Cameroun)', 'dialCode': 237, 'phoneFormat': '6 71 23 45 67'},
    'ca': {'code': 'ca', 'cn': '加拿大', 'en': 'Canada', 'dialCode': 1, 'phoneFormat': '(204) 234-5678'},
    'cv': {'code': 'cv', 'cn': '佛得角', 'en': 'Cape Verde (Kabu Verdi)', 'dialCode': 238, 'phoneFormat': '991 12 34'},
    'bq': {'code': 'bq', 'cn': '荷兰加勒比区', 'en': 'Caribbean Netherlands', 'dialCode': 599, 'phoneFormat': '318 1234'},
    'ky': {'code': 'ky', 'cn': '开曼群岛', 'en': 'Cayman Islands', 'dialCode': 1345, 'phoneFormat': '(345) 323-1234'},
    'cf': {
      'code': 'cf',
      'cn': '中非',
      'en': 'Central African Republic (République centrafricaine)',
      'dialCode': 236,
      'phoneFormat': '70 01 23 45'
    },
    'td': {'code': 'td', 'cn': '乍得', 'en': 'Chad (Tchad)', 'dialCode': 235, 'phoneFormat': '63 01 23 45'},
    'cl': {'code': 'cl', 'cn': '智利', 'en': 'Chile', 'dialCode': 56, 'phoneFormat': '09 6123 4567'},
    'cn': {'code': 'cn', 'cn': '中国', 'en': 'China (中国)', 'dialCode': 86, 'phoneFormat': '131 2345 6789'},
    'cx': {'code': 'cx', 'cn': '圣诞岛', 'en': 'Christmas Island', 'dialCode': 61, 'phoneFormat': '0412 345 678'},
    'cc': {'code': 'cc', 'cn': '科科斯群岛', 'en': 'Cocos (Keeling) Islands', 'dialCode': 61, 'phoneFormat': '0412 345 678'},
    'co': {'code': 'co', 'cn': '哥伦比亚', 'en': 'Colombia', 'dialCode': 57, 'phoneFormat': '321 1234567'},
    'km': {'code': 'km', 'cn': '科摩罗', 'en': 'Comoros (‫جزر القمر‬‎)', 'dialCode': 269, 'phoneFormat': '321 23 45'},
    // 'cd': {
    //   'code': 'cd',
    //   'cn': '',
    //   'name': 'Congo (DRC) (Jamhuri ya Kidemokrasia ya Kongo)',
    //   'dialCode': 243,
    //   'phoneFormat': '0991 234 567'
    // },
    'cg': {
      'code': 'cg',
      'cn': '刚果（布）',
      'en': 'Congo (Republic) (Congo-Brazzaville)',
      'dialCode': 242,
      'phoneFormat': '06 123 4567'
    },
    'ck': {'code': 'ck', 'cn': '库克群岛', 'en': 'Cook Islands', 'dialCode': 682, 'phoneFormat': '71 234'},
    'cr': {'code': 'cr', 'cn': '哥斯达黎加', 'en': 'Costa Rica', 'dialCode': 506, 'phoneFormat': '8312 3456'},
    'ci': {'code': 'ci', 'cn': '科特迪瓦', 'en': 'Côte d’Ivoire', 'dialCode': 225, 'phoneFormat': '01 23 45 67'},
    'hr': {'code': 'hr', 'cn': '克罗地亚', 'en': 'Croatia (Hrvatska)', 'dialCode': 385, 'phoneFormat': '091 234 5678'},
    // 'cu': {'code': 'cu', 'cn': '', 'en': 'Cuba', 'dialCode': 53, 'phoneFormat': '05 1234567'},
    'cw': {'code': 'cw', 'cn': '库拉索', 'en': 'Curaçao', 'dialCode': 599, 'phoneFormat': '9 518 1234'},
    'cy': {'code': 'cy', 'cn': '塞浦路斯', 'en': 'Cyprus (Κύπρος)', 'dialCode': 357, 'phoneFormat': '96 123456'},
    'cz': {'code': 'cz', 'cn': '捷克', 'en': 'Czech Republic (Česká republika)', 'dialCode': 420, 'phoneFormat': '601 123 456'},
    'dk': {'code': 'dk', 'cn': '丹麦', 'en': 'Denmark (Danmark)', 'dialCode': 45, 'phoneFormat': '20 12 34 56'},
    'dj': {'code': 'dj', 'cn': '吉布提', 'en': 'Djibouti', 'dialCode': 253, 'phoneFormat': '77 83 10 01'},
    'dm': {'code': 'dm', 'cn': '多米尼克', 'en': 'Dominica', 'dialCode': 1767, 'phoneFormat': '(767) 225-1234'},
    'do': {
      'code': 'do',
      'cn': '多米尼加',
      'en': 'Dominican Republic (República Dominicana)',
      'dialCode': 1,
      'phoneFormat': '(809) 234-5678'
    },
    'ec': {'code': 'ec', 'cn': '厄瓜多尔', 'en': 'Ecuador', 'dialCode': 593, 'phoneFormat': '099 123 4567'},
    'eg': {'code': 'eg', 'cn': '埃及', 'en': 'Egypt (‫مصر‬‎)', 'dialCode': 20, 'phoneFormat': '0100 123 4567'},
    'sv': {'code': 'sv', 'cn': '萨尔瓦多', 'en': 'El Salvador', 'dialCode': 503, 'phoneFormat': '7012 3456'},
    'gq': {
      'code': 'gq',
      'cn': '赤道几内亚',
      'en': 'Equatorial Guinea (Guinea Ecuatorial)',
      'dialCode': 240,
      'phoneFormat': '222 123 456'
    },
    'er': {'code': 'er', 'cn': '厄立特里亚', 'en': 'Eritrea', 'dialCode': 291, 'phoneFormat': '07 123 456'},
    'ee': {'code': 'ee', 'cn': '爱沙尼亚', 'en': 'Estonia (Eesti)', 'dialCode': 372, 'phoneFormat': '5123 4567'},
    'et': {'code': 'et', 'cn': '埃塞俄比亚', 'en': 'Ethiopia', 'dialCode': 251, 'phoneFormat': '091 123 4567'},
    'fk': {'code': 'fk', 'cn': '马尔维纳斯群岛（福克兰）', 'en': 'Falkland Islands (Islas Malvinas)', 'dialCode': 500, 'phoneFormat': '51234'},
    'fo': {'code': 'fo', 'cn': '法罗群岛', 'en': 'Faroe Islands (Føroyar)', 'dialCode': 298, 'phoneFormat': '211234'},
    'fj': {'code': 'fj', 'cn': '斐济群岛', 'en': 'Fiji', 'dialCode': 679, 'phoneFormat': '701 2345'},
    'fi': {'code': 'fi', 'cn': '芬兰', 'en': 'Finland (Suomi)', 'dialCode': 358, 'phoneFormat': '041 2345678'},
    'fr': {'code': 'fr', 'cn': '法国', 'en': 'France', 'dialCode': 33, 'phoneFormat': '06 12 34 56 78'},
    'gf': {'code': 'gf', 'cn': '法属圭亚那', 'en': 'French Guiana (Guyane française)', 'dialCode': 594, 'phoneFormat': '0694 20 12 34'},
    'pf': {
      'code': 'pf',
      'cn': '法属波利尼西亚',
      'en': 'French Polynesia (Polynésie française)',
      'dialCode': 689,
      'phoneFormat': '87 12 34 56'
    },
    'ga': {'code': 'ga', 'cn': '加蓬', 'en': 'Gabon', 'dialCode': 241, 'phoneFormat': '06 03 12 34'},
    'gm': {'code': 'gm', 'cn': '冈比亚', 'en': 'Gambia', 'dialCode': 220, 'phoneFormat': '301 2345'},
    'ge': {'code': 'ge', 'cn': '格鲁吉亚', 'en': 'Georgia (საქართველო)', 'dialCode': 995, 'phoneFormat': '555 12 34 56'},
    'de': {'code': 'de', 'cn': '德国', 'en': 'Germany (Deutschland)', 'dialCode': 49, 'phoneFormat': '01512 3456789'},
    'gh': {'code': 'gh', 'cn': '加纳', 'en': 'Ghana (Gaana)', 'dialCode': 233, 'phoneFormat': '023 123 4567'},
    'gi': {'code': 'gi', 'cn': '直布罗陀', 'en': 'Gibraltar', 'dialCode': 350, 'phoneFormat': '57123456'},
    'gr': {'code': 'gr', 'cn': '希腊', 'en': 'Greece (Ελλάδα)', 'dialCode': 30, 'phoneFormat': '691 234 5678'},
    'gl': {'code': 'gl', 'cn': '格陵兰', 'en': 'Greenland (Kalaallit Nunaat)', 'dialCode': 299, 'phoneFormat': '22 12 34'},
    'gd': {'code': 'gd', 'cn': '格林纳达', 'en': 'Grenada', 'dialCode': 1473, 'phoneFormat': '(473) 403-1234'},
    'gp': {'code': 'gp', 'cn': '瓜德罗普', 'en': 'Guadeloupe', 'dialCode': 590, 'phoneFormat': '0690 30-1234'},
    'gu': {'code': 'gu', 'cn': '关岛', 'en': 'Guam', 'dialCode': 1671, 'phoneFormat': '(671) 300-1234'},
    'gt': {'code': 'gt', 'cn': '危地马拉', 'en': 'Guatemala', 'dialCode': 502, 'phoneFormat': '5123 4567'},
    'gg': {'code': 'gg', 'cn': '根西岛', 'en': 'Guernsey', 'dialCode': 44, 'phoneFormat': '07781 123456'},
    'gn': {'code': 'gn', 'cn': '几内亚', 'en': 'Guinea (Guinée)', 'dialCode': 224, 'phoneFormat': '601 12 34 56'},
    'gw': {'code': 'gw', 'cn': '几内亚比绍', 'en': 'Guinea-Bissau (Guiné Bissau)', 'dialCode': 245, 'phoneFormat': '955 012 345'},
    'gy': {'code': 'gy', 'cn': '圭亚那', 'en': 'Guyana', 'dialCode': 592, 'phoneFormat': '609 1234'},
    'ht': {'code': 'ht', 'cn': '海地', 'en': 'Haiti', 'dialCode': 509, 'phoneFormat': '34 10 1234'},
    'hn': {'code': 'hn', 'cn': '洪都拉斯', 'en': 'Honduras', 'dialCode': 504, 'phoneFormat': '9123-4567'},
    'hk': {'code': 'hk', 'cn': '香港', 'en': 'Hong Kong (香港)', 'dialCode': 852, 'phoneFormat': '5123 4567'},
    'hu': {'code': 'hu', 'cn': '匈牙利', 'en': 'Hungary (Magyarország)', 'dialCode': 36, 'phoneFormat': '(20) 123 4567'},
    'is': {'code': 'is', 'cn': '冰岛', 'en': 'Iceland (Ísland)', 'dialCode': 354, 'phoneFormat': '611 1234'},
    'in': {'code': 'in', 'cn': '印度', 'en': 'India (भारत)', 'dialCode': 91, 'phoneFormat': '099876 54321'},
    'id': {'code': 'id', 'cn': '印尼', 'en': 'Indonesia', 'dialCode': 62, 'phoneFormat': '0812-345-678'},
    // 'ir': {'code': 'ir', 'cn': '', 'en': 'Iran (‫ایران‬‎)', 'dialCode': 98, 'phoneFormat': '0912 345 6789'},
    // 'iq': {'code': 'iq', 'cn': '', 'en': 'Iraq (‫العراق‬‎)', 'dialCode': 964, 'phoneFormat': '0791 234 5678'},
    'ie': {'code': 'ie', 'cn': '爱尔兰', 'en': 'Ireland', 'dialCode': 353, 'phoneFormat': '085 012 3456'},
    'im': {'code': 'im', 'cn': '马恩岛', 'en': 'Isle of Man', 'dialCode': 44, 'phoneFormat': '07924 123456'},
    'il': {'code': 'il', 'cn': '以色列', 'en': 'Israel (‫ישראל‬‎)', 'dialCode': 972, 'phoneFormat': '050-123-4567'},
    'it': {'code': 'it', 'cn': '意大利', 'en': 'Italy (Italia)', 'dialCode': 39, 'phoneFormat': '312 345 6789'},
    'jm': {'code': 'jm', 'cn': '牙买加', 'en': 'Jamaica', 'dialCode': 1876, 'phoneFormat': '(876) 210-1234'},
    'jp': {'code': 'jp', 'cn': '日本', 'en': 'Japan (日本)', 'dialCode': 81, 'phoneFormat': '090-1234-5678'},
    'je': {'code': 'je', 'cn': '泽西岛', 'en': 'Jersey', 'dialCode': 44, 'phoneFormat': '07797 123456'},
    'jo': {'code': 'jo', 'cn': '约旦', 'en': 'Jordan (‫الأردن‬‎)', 'dialCode': 962, 'phoneFormat': '07 9012 3456'},
    'kz': {'code': 'kz', 'cn': '哈萨克斯坦', 'en': 'Kazakhstan (Казахстан)', 'dialCode': 7, 'phoneFormat': '8 (771) 000 9998'},
    'ke': {'code': 'ke', 'cn': '肯尼亚', 'en': 'Kenya', 'dialCode': 254, 'phoneFormat': '0712 123456'},
    'ki': {'code': 'ki', 'cn': '基里巴斯', 'en': 'Kiribati', 'dialCode': 686, 'phoneFormat': '72012345'},
    'xk': {'code': 'xk', 'cn': '科索沃', 'en': 'Kosovo', 'dialCode': 383, 'phoneFormat': ''},
    'kw': {'code': 'kw', 'cn': '科威特', 'en': 'Kuwait (‫الكويت‬‎)', 'dialCode': 965, 'phoneFormat': '500 12345'},
    'kg': {'code': 'kg', 'cn': '吉尔吉斯斯坦', 'en': 'Kyrgyzstan (Кыргызстан)', 'dialCode': 996, 'phoneFormat': '0700 123 456'},
    'la': {'code': 'la', 'cn': '老挝', 'en': 'Laos (ລາວ)', 'dialCode': 856, 'phoneFormat': '020 23 123 456'},
    'lv': {'code': 'lv', 'cn': '拉脱维亚', 'en': 'Latvia (Latvija)', 'dialCode': 371, 'phoneFormat': '21 234 567'},
    'lb': {'code': 'lb', 'cn': '黎巴嫩', 'en': 'Lebanon (‫لبنان‬‎)', 'dialCode': 961, 'phoneFormat': '71 123 456'},
    'ls': {'code': 'ls', 'cn': '莱索托', 'en': 'Lesotho', 'dialCode': 266, 'phoneFormat': '5012 3456'},
    'lr': {'code': 'lr', 'cn': '利比里亚', 'en': 'Liberia', 'dialCode': 231, 'phoneFormat': '077 012 3456'},
    'ly': {'code': 'ly', 'cn': '利比亚', 'en': 'Libya (‫ليبيا‬‎)', 'dialCode': 218, 'phoneFormat': '091-2345678'},
    'li': {'code': 'li', 'cn': '列支敦士登', 'en': 'Liechtenstein', 'dialCode': 423, 'phoneFormat': '660 234 567'},
    'lt': {'code': 'lt', 'cn': '立陶宛', 'en': 'Lithuania (Lietuva)', 'dialCode': 370, 'phoneFormat': '(8-612) 34567'},
    'lu': {'code': 'lu', 'cn': '卢森堡', 'en': 'Luxembourg', 'dialCode': 352, 'phoneFormat': '628 123 456'},
    'mo': {'code': 'mo', 'cn': '澳门', 'en': 'Macau (澳門)', 'dialCode': 853, 'phoneFormat': '6612 3456'},
    'mk': {'code': 'mk', 'cn': '马其顿', 'en': 'Macedonia (FYROM) (Македонија)', 'dialCode': 389, 'phoneFormat': '072 345 678'},
    'mg': {'code': 'mg', 'cn': '马达加斯加', 'en': 'Madagascar (Madagasikara)', 'dialCode': 261, 'phoneFormat': '032 12 345 67'},
    'mw': {'code': 'mw', 'cn': '马拉维', 'en': 'Malawi', 'dialCode': 265, 'phoneFormat': '0991 23 45 67'},
    'my': {'code': 'my', 'cn': '马来西亚', 'en': 'Malaysia', 'dialCode': 60, 'phoneFormat': '012-345 6789'},
    'mv': {'code': 'mv', 'cn': '马尔代夫', 'en': 'Maldives', 'dialCode': 960, 'phoneFormat': '771-2345'},
    'ml': {'code': 'ml', 'cn': '马里', 'en': 'Mali', 'dialCode': 223, 'phoneFormat': '65 01 23 45'},
    'mt': {'code': 'mt', 'cn': '马耳他', 'en': 'Malta', 'dialCode': 356, 'phoneFormat': '9696 1234'},
    'mh': {'code': 'mh', 'cn': '马绍尔群岛', 'en': 'Marshall Islands', 'dialCode': 692, 'phoneFormat': '235-1234'},
    'mq': {'code': 'mq', 'cn': '马提尼克', 'en': 'Martinique', 'dialCode': 596, 'phoneFormat': '0696 20 12 34'},
    'mr': {'code': 'mr', 'cn': '毛里塔尼亚', 'en': 'Mauritania (‫موريتانيا‬‎)', 'dialCode': 222, 'phoneFormat': '22 12 34 56'},
    'mu': {'code': 'mu', 'cn': '毛里求斯', 'en': 'Mauritius (Moris)', 'dialCode': 230, 'phoneFormat': '5251 2345'},
    'yt': {'code': 'yt', 'cn': '马约特', 'en': 'Mayotte', 'dialCode': 262, 'phoneFormat': '0639 12 34 56'},
    'mx': {'code': 'mx', 'cn': '墨西哥', 'en': 'Mexico (México)', 'dialCode': 52, 'phoneFormat': '044 222 123 4567'},
    'fm': {'code': 'fm', 'cn': '密克罗尼西亚联邦', 'en': 'Micronesia', 'dialCode': 691, 'phoneFormat': '350 1234'},
    'md': {'code': 'md', 'cn': '摩尔多瓦', 'en': 'Moldova (Republica Moldova)', 'dialCode': 373, 'phoneFormat': '0621 12 345'},
    'mc': {'code': 'mc', 'cn': '摩纳哥', 'en': 'Monaco', 'dialCode': 377, 'phoneFormat': '06 12 34 56 78'},
    'mn': {'code': 'mn', 'cn': '蒙古国', 'en': 'Mongolia (Монгол)', 'dialCode': 976, 'phoneFormat': '8812 3456'},
    'me': {'code': 'me', 'cn': '黑山', 'en': 'Montenegro (Crna Gora)', 'dialCode': 382, 'phoneFormat': '067 622 901'},
    'ms': {'code': 'ms', 'cn': '蒙塞拉特岛', 'en': 'Montserrat', 'dialCode': 1664, 'phoneFormat': '(664) 492-3456'},
    'ma': {'code': 'ma', 'cn': '摩洛哥', 'en': 'Morocco (‫المغرب‬‎)', 'dialCode': 212, 'phoneFormat': '0650-123456'},
    'mz': {'code': 'mz', 'cn': '莫桑比克', 'en': 'Mozambique (Moçambique)', 'dialCode': 258, 'phoneFormat': '82 123 4567'},
    'mm': {'code': 'mm', 'cn': '缅甸', 'en': 'Myanmar (Burma) (မြန်မာ)', 'dialCode': 95, 'phoneFormat': '09 212 3456'},
    'na': {'code': 'na', 'cn': '纳米比亚', 'en': 'Namibia (Namibië)', 'dialCode': 264, 'phoneFormat': '081 123 4567'},
    'nr': {'code': 'nr', 'cn': '瑙鲁', 'en': 'Nauru', 'dialCode': 674, 'phoneFormat': '555 1234'},
    'np': {'code': 'np', 'cn': '尼泊尔', 'en': 'Nepal (नेपाल)', 'dialCode': 977, 'phoneFormat': '984-1234567'},
    'nl': {'code': 'nl', 'cn': '荷兰', 'en': 'Netherlands (Nederland)', 'dialCode': 31, 'phoneFormat': '06 12345678'},
    'nc': {'code': 'nc', 'cn': '新喀里多尼亚', 'en': 'New Caledonia (Nouvelle-Calédonie)', 'dialCode': 687, 'phoneFormat': '75.12.34'},
    'nz': {'code': 'nz', 'cn': '新西兰', 'en': 'New Zealand', 'dialCode': 64, 'phoneFormat': '021 123 4567'},
    'ni': {'code': 'ni', 'cn': '尼加拉瓜', 'en': 'Nicaragua', 'dialCode': 505, 'phoneFormat': '8123 4567'},
    'ne': {'code': 'ne', 'cn': '尼日尔', 'en': 'Niger (Nijar)', 'dialCode': 227, 'phoneFormat': '93 12 34 56'},
    'ng': {'code': 'ng', 'cn': '尼日利亚', 'en': 'Nigeria', 'dialCode': 234, 'phoneFormat': '0802 123 4567'},
    'nu': {'code': 'nu', 'cn': '纽埃', 'en': 'Niue', 'dialCode': 683, 'phoneFormat': '1234'},
    'nf': {'code': 'nf', 'cn': '诺福克岛', 'en': 'Norfolk Island', 'dialCode': 672, 'phoneFormat': '3 81234'},
    // 'kp': {'code': 'kp', 'cn': '', 'en': 'North Korea (조선 민주주의 인민 공화국)', 'dialCode': 850, 'phoneFormat': '0192 123 4567'},
    'mp': {'code': 'mp', 'cn': '北马里亚纳群岛', 'en': 'Northern Mariana Islands', 'dialCode': 1670, 'phoneFormat': '(670) 234-5678'},
    'no': {'code': 'no', 'cn': '挪威', 'en': 'Norway (Norge)', 'dialCode': 47, 'phoneFormat': '406 12 345'},
    'om': {'code': 'om', 'cn': '阿曼', 'en': 'Oman (‫عُمان‬‎)', 'dialCode': 968, 'phoneFormat': '9212 3456'},
    'pk': {'code': 'pk', 'cn': '巴基斯坦', 'en': 'Pakistan (‫پاکستان‬‎)', 'dialCode': 92, 'phoneFormat': '0301 2345678'},
    'pw': {'code': 'pw', 'cn': '帕劳', 'en': 'Palau', 'dialCode': 680, 'phoneFormat': '620 1234'},
    'ps': {'code': 'ps', 'cn': '巴勒斯坦', 'en': 'Palestine (‫فلسطين‬‎)', 'dialCode': 970, 'phoneFormat': '0599 123 456'},
    'pa': {'code': 'pa', 'cn': '巴拿马', 'en': 'Panama (Panamá)', 'dialCode': 507, 'phoneFormat': '6001-2345'},
    'pg': {'code': 'pg', 'cn': '巴布亚新几内亚', 'en': 'Papua New Guinea', 'dialCode': 675, 'phoneFormat': '681 2345'},
    'py': {'code': 'py', 'cn': '巴拉圭', 'en': 'Paraguay', 'dialCode': 595, 'phoneFormat': '0961 456789'},
    'pe': {'code': 'pe', 'cn': '秘鲁', 'en': 'Peru (Perú)', 'dialCode': 51, 'phoneFormat': '912 345 678'},
    'ph': {'code': 'ph', 'cn': '菲律宾', 'en': 'Philippines', 'dialCode': 63, 'phoneFormat': '0905 123 4567'},
    'pl': {'code': 'pl', 'cn': '波兰', 'en': 'Poland (Polska)', 'dialCode': 48, 'phoneFormat': '512 345 678'},
    'pt': {'code': 'pt', 'cn': '葡萄牙', 'en': 'Portugal', 'dialCode': 351, 'phoneFormat': '912 345 678'},
    'pr': {'code': 'pr', 'cn': '波多黎各', 'en': 'Puerto Rico', 'dialCode': 1, 'phoneFormat': '(787) 234-5678'},
    'qa': {'code': 'qa', 'cn': '卡塔尔', 'en': 'Qatar (‫قطر‬‎)', 'dialCode': 974, 'phoneFormat': '3312 3456'},
    're': {'code': 're', 'cn': '留尼汪', 'en': 'Réunion (La Réunion)', 'dialCode': 262, 'phoneFormat': '0692 12 34 56'},
    'ro': {'code': 'ro', 'cn': '罗马尼亚', 'en': 'Romania (România)', 'dialCode': 40, 'phoneFormat': '0712 345 678'},
    'ru': {'code': 'ru', 'cn': '俄罗斯', 'en': 'Russia (Россия)', 'dialCode': 7, 'phoneFormat': '8 (912) 345-67-89'},
    'rw': {'code': 'rw', 'cn': '卢旺达', 'en': 'Rwanda', 'dialCode': 250, 'phoneFormat': '0720 123 456'},
    'bl': {
      'code': 'bl',
      'cn': '圣巴泰勒米岛',
      'en': 'Saint Barthélemy (Saint-Barthélemy)',
      'dialCode': 590,
      'phoneFormat': '0690 30-1234'
    },
    'sh': {'code': 'sh', 'cn': '圣赫勒拿', 'en': 'Saint Helena', 'dialCode': 290, 'phoneFormat': '51234'},
    'kn': {'code': 'kn', 'cn': '圣基茨和尼维斯', 'en': 'Saint Kitts and Nevis', 'dialCode': 1869, 'phoneFormat': '(869) 765-2917'},
    'lc': {'code': 'lc', 'cn': '圣卢西亚', 'en': 'Saint Lucia', 'dialCode': 1758, 'phoneFormat': '(758) 284-5678'},
    'mf': {
      'code': 'mf',
      "cn": "法属圣马丁",
      'en': 'Saint Martin (Saint-Martin (partie française))',
      'dialCode': 590,
      'phoneFormat': '0690 30-1234'
    },
    'pm': {
      'code': 'pm',
      "cn": "圣皮埃尔和密克隆",
      'en': 'Saint Pierre and Miquelon (Saint-Pierre-et-Miquelon)',
      'dialCode': 508,
      'phoneFormat': '055 12 34'
    },
    'vc': {
      'code': 'vc',
      "cn": "圣文森特和格林纳丁斯",
      'en': 'Saint Vincent and the Grenadines',
      'dialCode': 1784,
      'phoneFormat': '(784) 430-1234'
    },
    'ws': {'code': 'ws', 'cn': '萨摩亚', 'en': 'Samoa', 'dialCode': 685, 'phoneFormat': '601234'},
    'sm': {'code': 'sm', 'cn': '圣马力诺', 'en': 'San Marino', 'dialCode': 378, 'phoneFormat': '66 66 12 12'},
    'st': {
      'code': 'st',
      "cn": "圣多美和普林西比",
      'en': 'São Tomé and Príncipe (São Tomé e Príncipe)',
      'dialCode': 239,
      'phoneFormat': '981 2345'
    },
    'sa': {
      'code': 'sa',
      "cn": "沙特阿拉伯",
      'en': 'Saudi Arabia (‫المملكة العربية السعودية‬‎)',
      'dialCode': 966,
      'phoneFormat': '051 234 5678'
    },
    'sn': {'code': 'sn', 'cn': '塞内加尔', 'en': 'Senegal (Sénégal)', 'dialCode': 221, 'phoneFormat': '70 123 45 67'},
    'rs': {'code': 'rs', 'cn': '塞尔维亚', 'en': 'Serbia (Србија)', 'dialCode': 381, 'phoneFormat': '060 1234567'},
    'sc': {'code': 'sc', 'cn': '塞舌尔', 'en': 'Seychelles', 'dialCode': 248, 'phoneFormat': '2 510 123'},
    'sl': {'code': 'sl', 'cn': '塞拉利昂', 'en': 'Sierra Leone', 'dialCode': 232, 'phoneFormat': '(025) 123456'},
    'sg': {'code': 'sg', 'cn': '新加坡', 'en': 'Singapore', 'dialCode': 65, 'phoneFormat': '8123 4567'},
    'sx': {'code': 'sx', 'cn': '荷属圣马丁', 'en': 'Sint Maarten', 'dialCode': 1721, 'phoneFormat': '(721) 520-5678'},
    'sk': {'code': 'sk', 'cn': '斯洛伐克', 'en': 'Slovakia (Slovensko)', 'dialCode': 421, 'phoneFormat': '0912 123 456'},
    'si': {'code': 'si', 'cn': '斯洛文尼亚', 'en': 'Slovenia (Slovenija)', 'dialCode': 386, 'phoneFormat': '031 234 567'},
    'sb': {'code': 'sb', 'cn': '所罗门群岛', 'en': 'Solomon Islands', 'dialCode': 677, 'phoneFormat': '74 21234'},
    'so': {'code': 'so', 'cn': '索马里', 'en': 'Somalia (Soomaaliya)', 'dialCode': 252, 'phoneFormat': '7 1123456'},
    'za': {'code': 'za', 'cn': '南非', 'en': 'South Africa', 'dialCode': 27, 'phoneFormat': '071 123 4567'},
    'kr': {'code': 'kr', 'cn': '韩国', 'en': 'South Korea (대한민국)', 'dialCode': 82, 'phoneFormat': '010-0000-0000'},
    'ss': {'code': 'ss', 'cn': '南苏丹', 'en': 'South Sudan (‫جنوب السودان‬‎)', 'dialCode': 211, 'phoneFormat': '0977 123 456'},
    'es': {'code': 'es', 'cn': '西班牙', 'en': 'Spain (España)', 'dialCode': 34, 'phoneFormat': '612 34 56 78'},
    'lk': {'code': 'lk', 'cn': '斯里兰卡', 'en': 'Sri Lanka (ශ්‍රී ලංකාව)', 'dialCode': 94, 'phoneFormat': '071 234 5678'},
    // 'sd': {'code': 'sd', 'cn': '', 'en': 'Sudan (‫السودان‬‎)', 'dialCode': 249, 'phoneFormat': '091 123 1234'},
    'sr': {'code': 'sr', 'cn': '苏里南', 'en': 'Suriname', 'dialCode': 597, 'phoneFormat': '741-2345'},
    'sj': {'code': 'sj', 'cn': '斯瓦尔巴群岛和扬马延岛', 'en': 'Svalbard and Jan Mayen', 'dialCode': 47, 'phoneFormat': '412 34 567'},
    'sz': {'code': 'sz', 'cn': '斯威士兰', 'en': 'Swaziland', 'dialCode': 268, 'phoneFormat': '7612 3456'},
    'se': {'code': 'se', 'cn': '瑞典', 'en': 'Sweden (Sverige)', 'dialCode': 46, 'phoneFormat': '070-123 45 67'},
    'ch': {'code': 'ch', 'cn': '瑞士', 'en': 'Switzerland (Schweiz)', 'dialCode': 41, 'phoneFormat': '078 123 45 67'},
    // 'sy': {'code': 'sy', 'cn': '', 'en': 'Syria (‫سوريا‬‎)', 'dialCode': 963, 'phoneFormat': '0944 567 890'},
    'tw': {'code': 'tw', 'cn': '台湾', 'en': 'Taiwan (台灣)', 'dialCode': 886, 'phoneFormat': '0912 345 678'},
    'tj': {'code': 'tj', 'cn': '塔吉克斯坦', 'en': 'Tajikistan', 'dialCode': 992, 'phoneFormat': '(8) 917 12 3456'},
    'tz': {'code': 'tz', 'cn': '坦桑尼亚', 'en': 'Tanzania', 'dialCode': 255, 'phoneFormat': '0621 234 567'},
    'th': {'code': 'th', 'cn': '泰国', 'en': 'Thailand (ไทย)', 'dialCode': 66, 'phoneFormat': '081 234 5678'},
    'tl': {'code': 'tl', 'cn': '东帝汶', 'en': 'Timor-Leste', 'dialCode': 670, 'phoneFormat': '7721 2345'},
    'tg': {'code': 'tg', 'cn': '多哥', 'en': 'Togo', 'dialCode': 228, 'phoneFormat': '90 11 23 45'},
    'tk': {'code': 'tk', 'cn': '托克劳', 'en': 'Tokelau', 'dialCode': 690, 'phoneFormat': '7290'},
    'to': {'code': 'to', 'cn': '汤加', 'en': 'Tonga', 'dialCode': 676, 'phoneFormat': '771 5123'},
    'tt': {'code': 'tt', 'cn': '特立尼达和多巴哥', 'en': 'Trinidad and Tobago', 'dialCode': 1868, 'phoneFormat': '(868) 291-1234'},
    'tn': {'code': 'tn', 'cn': '突尼斯', 'en': 'Tunisia (‫تونس‬‎)', 'dialCode': 216, 'phoneFormat': '20 123 456'},
    'tr': {'code': 'tr', 'cn': '土耳其', 'en': 'Turkey (Türkiye)', 'dialCode': 90, 'phoneFormat': '0501 234 56 78'},
    'tm': {'code': 'tm', 'cn': '土库曼斯坦', 'en': 'Turkmenistan', 'dialCode': 993, 'phoneFormat': '8 66 123456'},
    'tc': {'code': 'tc', 'cn': '特克斯和凯科斯群岛', 'en': 'Turks and Caicos Islands', 'dialCode': 1649, 'phoneFormat': '(649) 231-1234'},
    'tv': {'code': 'tv', 'cn': '图瓦卢', 'en': 'Tuvalu', 'dialCode': 688, 'phoneFormat': '901234'},
    'us': {'code': 'us', 'cn': '美国本土外小岛屿', 'en': 'United States', 'dialCode': 1, 'phoneFormat': '(201) 555-0123'},
    'gb': {'code': 'gb', 'cn': '英国', 'en': 'United Kingdom', 'dialCode': 44, 'phoneFormat': '07400 123456'},
    'vi': {'code': 'vi', 'cn': '美属维尔京群岛', 'en': 'U.S. Virgin Islands', 'dialCode': 1340, 'phoneFormat': '(340) 642-1234'},
    'ug': {'code': 'ug', 'cn': '乌干达', 'en': 'Uganda', 'dialCode': 256, 'phoneFormat': '0712 345678'},
    'ua': {'code': 'ua', 'cn': '乌克兰', 'en': 'Ukraine (Україна)', 'dialCode': 380, 'phoneFormat': '039 123 4567'},
    'ae': {
      'code': 'ae',
      'name': 'United Arab Emirates (‫الإمارات العربية المتحدة‬‎)',
      'dialCode': 971,
      'phoneFormat': '050 123 4567'
    },
    'uy': {'code': 'uy', 'name': 'Uruguay', 'dialCode': 598, 'phoneFormat': '094 231 234'},
    'uz': {'code': 'uz', 'name': 'Uzbekistan (Oʻzbekiston)', 'dialCode': 998, 'phoneFormat': '8 91 234 56 78'},
    'vu': {'code': 'vu', 'name': 'Vanuatu', 'dialCode': 678, 'phoneFormat': '591 2345'},
    'va': {'code': 'va', 'name': 'Vatican City (Città del Vaticano)', 'dialCode': 39, 'phoneFormat': '312 345 6789'},
    've': {'code': 've', 'name': 'Venezuela', 'dialCode': 58, 'phoneFormat': '0412-1234567'},
    'vn': {'code': 'vn', 'name': 'Vietnam (Việt Nam)', 'dialCode': 84, 'phoneFormat': '091 234 56 78'},
    'wf': {'code': 'wf', 'name': 'Wallis and Futuna', 'dialCode': 681, 'phoneFormat': '50 12 34'},
    'eh': {
      'code': 'eh',
      "cn": "阿联酋",
      'en': 'Western Sahara (‫الصحراء الغربية‬‎)',
      'dialCode': 212,
      'phoneFormat': '0650-123456'
    },
    'ye': {'code': 'ye', 'cn': '也门', 'en': 'Yemen (‫اليمن‬‎)', 'dialCode': 967, 'phoneFormat': '0712 345 678'},
    'zm': {'code': 'zm', 'cn': '赞比亚', 'en': 'Zambia', 'dialCode': 260, 'phoneFormat': '095 5123456'},
    // 'zw': {'code': 'zw', 'cn': '', 'en': 'Zimbabwe', 'dialCode': 263, 'phoneFormat': '071 123 4567'},
    'ax': {'code': 'ax', 'cn': '奥兰群岛', 'en': 'Åland Islands', 'dialCode': 358, 'phoneFormat': '041 2345678'}
  }
  hideSubMenu = true
  @Input() currentCode = ''
  @Input() currentAreaCode = ''
  @Output() setCurrentCode = new EventEmitter<any>();
  @Output() setCurrentAreaCode = new EventEmitter<any>();
  
  // currentCode = ''
  countriesCodes:string[] = []
  constructor() {
    this.countriesCodes = Object.keys(this.countries)
    this.currentCode = Object.keys(this.countries)[0]    
  }
  get currentData () {
    return this.countries[this.currentCode]
  }
  get frontCountriesArray () {
    const countries = []
    this.countriesCodes.forEach(code => {
      countries.push(this.countries[code])
    })
    return countries
  }
  get lang () {
    const lang = localStorage.getItem('i18n')
    if (lang && lang === 'zh-CN') {
      return 'cn'
    } else {
      return 'en'
    }
  }
  cutFlag (item) {
    // this.currentCode = item.code;
    this.setCurrentCode.emit(item.code);
    this.setCurrentAreaCode.emit(item.dialCode);
 
    this.hideSubMenu = true; 
    this.setCountry(item); 
    // this.currentAreaCode=item.dialCode;
  }
  setCountry (item) {}
  ngOnInit() { }

}