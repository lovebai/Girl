package dao

import "Girl/model"

//查询
func (mgr *manager) GetWebSiteInfo() int64 {
	info := model.Siteinfo{}
	return mgr.db.First(&info).RowsAffected
}

//网站配置
func (mgr *manager) CreateSiteinfo() int64 {
	siteinfo := model.Siteinfo{
		OptionId:       1,
		SiteName:       "Like_Girl for Golang",
		SiteLogo:       "Like_Girl_Golang",
		Keywords:       "Like_Girl for Golang,情侣小站,开源情侣网站,Golang情侣网站,情侣记录,情侣网站,情侣项目,情侣小窝,Love,LikeGirl,Golang情侣小站,情侣小站使用教程,情侣小站使用文档",
		Discription:    "Like_Girl for Golang 是一款使用Golang语言开发的网页情侣小站。",
		SiteIcon:       "/static/img/favicon.ico",
		SiteBlur:       0,
		SitePajx:       0,
		Watch:          0,
		BoyName:        "Ki",
		GirlName:       "Li",
		StartTime:      1718121600000,
		BgimgUrl:       "/api/RandImg",
		CardOne:        "点点滴滴",
		CardTwo:        "留言板",
		CardThree:      "关于我们",
		SiteIcp:        "渝ICP备0000000号",
		SiteGaIcp:      "渝公网安备50000002000110号",
		SiteCopyright:  "Copyright © 2022 - 2024 Like_Girl_Golang All Rights Reserved.",
		SiteDesc:       "喜欢花 喜欢浪漫 喜欢你.~",
		CardOneDesc:    "有人愿意听你碎碎念念也很浪漫",
		CardTwoDesc:    "在这里写下我们的留言祝福1",
		CardThreeDesc:  "我们之间认识的经历回忆",
		LenvingSum:     25,
		SiteHeadStyle:  "<style>\n .avatar-aa {\n            width: 3em;\n            height: 3em;\n            border-radius: 50%;\n            box-shadow: 0 2px 20px #c5c5c575;\n            border: 2px solid #fff;\n            margin-right: 0.8rem;\n        }  \n</style>",
		SiteRegexp:     "`~!@#$^&*()=|{}':;',\\\\[\\\\].<>/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？",
		SiteFootJs:     "<script>\nlet a = \"hello\"\nconsole.log(a);\n</script>",
		BlockWord:      "sb|傻逼",
		LinkType:       0,
		LinkApiUrl:     "",
		LinkApiToken:   "",
		AvaterShowType: 0,
		BoyImgUrl:      "",
		GirlImgUrl:     "",
	}
	return mgr.db.Create(&siteinfo).RowsAffected
}

//点滴
func (mgr *manager) CreateArticle(auname string) int64 {
	st := model.Article{
		ArticleId:    1,
		ArticleTitle: "Like_Girl 默认文章语法",
		ArticleContext: `<quote>引用内容样式</quote>

<hr>


<h1>H1文字大小演示</h1>

<hr>


<h2>H2文字大小演示</h2>

<hr>


<h3>H3文字大小演示</h3>

<hr>


<h4>H4文字大小演示</h4>

<hr>


<h5>H5文字大小演示</h5>

<hr>


<h6>H6文字大小演示</h6>

<hr>


<b>加粗字体</b>
<s>删除线字体</s>
<i>斜体</i>
<code>强调内容</code>

<center>文本居中</center>


<!--分割线-->
<hr>
<quote>插入图片</quote>
<img alt="" src="https://img.gejiba.com/images/b0d44ed67e25235f552aacbe32d81b5c.jpg">
<!--分割线-->
<hr>
<quote>插入视频</quote>
<iframe src="https://player.kikiw.cn/player/?url=https://www.kikiw.cn/love.mp4" allowfullscreen="true"></iframe>
<!--分割线-->
<hr>`,
		ArticleAuthor: auname,
		ArticleTime:   1718121600000}
	return mgr.db.Create(&st).RowsAffected
}

//留言
func (mgr *manager) CreateLenving() int64 {
	le := model.Lenving{
		LenvingId:   1,
		LenvingName: "Ki.",
		LenvingQq:   "3439780232",
		LenvingText: "LikeGirl#5.1.0默认留言",
		LenvingTime: 1718090394000,
		LenvingIp:   "123.123.123.123",
		LenvingCity: "北京市",
	}
	return mgr.db.Create(&le).RowsAffected
}

//相册
func (mgr *manager) CreatePhoto() int64 {
	ph := []model.Photo{
		{ImgId: 1, ImgText: "Like_Girl — 5.1.0相册测试", ImgUrl: "https://www.kikiw.cn/liuyi.png", ImgTime: 1718010394000},
		{ImgId: 2, ImgText: "夏至末至 我最喜欢听的追光者~", ImgUrl: "https://img.gejiba.com/images/43f3e4449c2fe004903d87164998b9e6.jpg", ImgTime: 1718090294000},
		{ImgId: 3, ImgText: "“真的很想和你去一趟浪漫的厦门”", ImgUrl: "https://img.gejiba.com/images/901a33ff29f3b9480fc8572856cf7664.jpg", ImgTime: 1718030394000},
		{ImgId: 4, ImgText: "“真的很想和你去一趟浪漫的厦门”", ImgUrl: "https://img.gejiba.com/images/dfafa31a39bc0775306bfd83c3571c68.jpg", ImgTime: 1718040394000},
		{ImgId: 5, ImgText: `追光者多年后听到还是会心头一颤`, ImgUrl: "https://img.gejiba.com/images/f18a5c0ecd51fc6cc43b5afef3e92837.jpg", ImgTime: 1718050394000},
		{ImgId: 6, ImgText: "太可爱啦 好像是叫“团团”来着~", ImgUrl: "https://img.gejiba.com/images/8344adb24e8440e83c0d2b36cba19fdd.jpg", ImgTime: 1718060394000},
		{ImgId: 7, ImgText: "最近不错的一部港片 很好看", ImgUrl: "https://img.gejiba.com/images/d929e9197da0237da2742adb19491abd.jpg", ImgTime: 1718070344000},
		{ImgId: 8, ImgText: "入手了一辆公路车 骑了30公里 现在屁股发麻...", ImgUrl: "https://img.gejiba.com/images/688553237cdbaacd58ec513e322c47ca.jpg", ImgTime: 1718080294000},
		{ImgId: 9, ImgText: "太好吃啦 这个牛蛙", ImgUrl: "https://img.gejiba.com/images/9b0940c31dfa5bdded0af45e7fa5b581.jpg", ImgTime: 1718095394000},
		{ImgId: 10, ImgText: "我们去草原吧", ImgUrl: "https://img.gejiba.com/images/06cac83c9839f0fa82876f6facf12bb0.jpg", ImgTime: 1718096394000},
	}
	return mgr.db.Create(&ph).RowsAffected
}

//清单
func (mgr *manager) CreateTodoList() int64 {
	tl := []model.TodoList{
		{
			ListId:     1,
			ListStatus: 1,
			ListText:   "一起去电影院看一场电影🎬",
			ListImgurl: "https://img.gejiba.com/images/d929e9197da0237da2742adb19491abd.jpg",
		},
		{
			ListId:     2,
			ListStatus: 0,
			ListText:   "一起穿情侣装逛街🧡",
			ListImgurl: "",
		},
		{
			ListId:     3,
			ListStatus: 0,
			ListText:   "一起去一趟迪士尼游乐园🍉",
			ListImgurl: "",
		},
		{
			ListId:     4,
			ListStatus: 0,
			ListText:   "一起去游泳🍇",
			ListImgurl: "",
		},
		{
			ListId:     5,
			ListStatus: 0,
			ListText:   "一起唱次歌并且录下来🍓",
			ListImgurl: "",
		},
		{
			ListId:     6,
			ListStatus: 0,
			ListText:   "一起在厨房做次饭🍈",
			ListImgurl: "",
		},
		{
			ListId:     7,
			ListStatus: 0,
			ListText:   "一起过次烛光晚餐🍒",
			ListImgurl: "",
		},
		{
			ListId:     8,
			ListStatus: 0,
			ListText:   "一起过生日🍑",
			ListImgurl: "",
		},
		{
			ListId:     9,
			ListStatus: 0,
			ListText:   "一起打扫卫生🥭",
			ListImgurl: "",
		},
		{
			ListId:     10,
			ListStatus: 0,
			ListText:   "一起给对方写信，然后读给对方听🍍",
			ListImgurl: "",
		},
		{
			ListId:     11,
			ListStatus: 0,
			ListText:   "一起去一次鬼屋🥥",
			ListImgurl: "",
		},
		{
			ListId:     12,
			ListStatus: 0,
			ListText:   "一起去蹦极🥝",
			ListImgurl: "",
		},
		{
			ListId:     13,
			ListStatus: 0,
			ListText:   "一起养一只宠物🍅",
			ListImgurl: "",
		},
		{
			ListId:     14,
			ListStatus: 0,
			ListText:   "一起研究口红色号🌽",
			ListImgurl: "",
		},
		{
			ListId:     15,
			ListStatus: 0,
			ListText:   "一起给对方化妆🧅",
			ListImgurl: "",
		},
		{
			ListId:     16,
			ListStatus: 0,
			ListText:   "一起为对方抹指甲油🥕",
			ListImgurl: "",
		},
		{
			ListId:     17,
			ListStatus: 0,
			ListText:   "一起去做次陶艺🥗",
			ListImgurl: "",
		},
		{
			ListId:     18,
			ListStatus: 0,
			ListText:   "一起去吃一次全家桶🥔",
			ListImgurl: "",
		},
		{
			ListId:     19,
			ListStatus: 0,
			ListText:   "一起熬夜通宵跨年🍠",
			ListImgurl: "",
		},
		{
			ListId:     20,
			ListStatus: 0,
			ListText:   "一起去旅游🥯",
			ListImgurl: "",
		},
		{
			ListId:     21,
			ListStatus: 0,
			ListText:   "一起去爬山⛰🧇",
			ListImgurl: "",
		},
		{
			ListId:     22,
			ListStatus: 0,
			ListText:   "一起坐一次摩天轮🥞",
			ListImgurl: "",
		},
		{
			ListId:     23,
			ListStatus: 0,
			ListText:   "一起拍视频记录生活🧀",
			ListImgurl: "",
		},
		{
			ListId:     24,
			ListStatus: 0,
			ListText:   "一起为对方刷牙，然后亲亲🍗",
			ListImgurl: "",
		},
		{
			ListId:     25,
			ListStatus: 0,
			ListText:   "一起去看一次海，去沙滩🍖",
			ListImgurl: "",
		},
		{
			ListId:     26,
			ListStatus: 0,
			ListText:   "互穿对方的衣服，拍照留念🥩",
			ListImgurl: "",
		},
		{
			ListId:     27,
			ListStatus: 0,
			ListText:   "一起逛超市买好吃的🍤",
			ListImgurl: "",
		},
		{
			ListId:     28,
			ListStatus: 0,
			ListText:   "一起坐一次热气球🥓",
			ListImgurl: "",
		},
		{
			ListId:     29,
			ListStatus: 0,
			ListText:   "一起看书，分享自己喜欢的书籍🍔",
			ListImgurl: "",
		},
		{
			ListId:     30,
			ListStatus: 0,
			ListText:   "一起在下雨天追剧🍟",
			ListImgurl: "",
		},
		{
			ListId:     31,
			ListStatus: 0,
			ListText:   "一起做一次蛋糕甜点🌭",
			ListImgurl: "",
		},
		{
			ListId:     32,
			ListStatus: 0,
			ListText:   "一起看日出看日落🍕",
			ListImgurl: "",
		},
		{
			ListId:     33,
			ListStatus: 0,
			ListText:   "一起上下班，坐地铁🍝",
			ListImgurl: "",
		},
		{
			ListId:     34,
			ListStatus: 0,
			ListText:   "一起坐一次飞机🥪",
			ListImgurl: "",
		},
		{
			ListId:     35,
			ListStatus: 0,
			ListText:   "一起种花草🌮",
			ListImgurl: "",
		},
		{
			ListId:     36,
			ListStatus: 0,
			ListText:   "一起用情侣手机壳🌯",
			ListImgurl: "",
		},
		{
			ListId:     37,
			ListStatus: 0,
			ListText:   "一起去一次海底世界🥙",
			ListImgurl: "",
		},
		{
			ListId:     38,
			ListStatus: 0,
			ListText:   "一起喝醉一次🧆",
			ListImgurl: "",
		},
		{
			ListId:     39,
			ListStatus: 0,
			ListText:   "一起打扑克牌🍜",
			ListImgurl: "",
		},
		{
			ListId:     40,
			ListStatus: 0,
			ListText:   "一起修理电器🍲",
			ListImgurl: "",
		},
		{
			ListId:     41,
			ListStatus: 0,
			ListText:   "一起看烟花🥘",
			ListImgurl: "",
		},
		{
			ListId:     42,
			ListStatus: 0,
			ListText:   "一起吃火锅🧂",
			ListImgurl: "",
		},
		{
			ListId:     43,
			ListStatus: 0,
			ListText:   "一起庆祝恋爱纪念日🧈",
			ListImgurl: "",
		},
		{
			ListId:     44,
			ListStatus: 0,
			ListText:   "一起看雪，堆雪人🍥",
			ListImgurl: "",
		},
		{
			ListId:     45,
			ListStatus: 0,
			ListText:   "一起和朋友们去吃饭🍱",
			ListImgurl: "",
		},
		{
			ListId:     46,
			ListStatus: 0,
			ListText:   "一起跳舞🍣",
			ListImgurl: "",
		},
		{
			ListId:     47,
			ListStatus: 0,
			ListText:   "一起听音乐，听同一首歌🍙",
			ListImgurl: "",
		},
		{
			ListId:     48,
			ListStatus: 0,
			ListText:   "一起坐一次船🍛",
			ListImgurl: "",
		},
		{
			ListId:     49,
			ListStatus: 0,
			ListText:   "一起露营，住一次帐篷🍘",
			ListImgurl: "",
		},
		{
			ListId:     50,
			ListStatus: 0,
			ListText:   "一起DIY手工🍚",
			ListImgurl: "",
		},
		{
			ListId:     51,
			ListStatus: 0,
			ListText:   "给对方准备礼物🥟",
			ListImgurl: "",
		},
		{
			ListId:     52,
			ListStatus: 0,
			ListText:   "一起去我们上过的小学，中学，大学🍢",
			ListImgurl: "",
		},
		{
			ListId:     53,
			ListStatus: 0,
			ListText:   "一起在沙发上躺着🍡",
			ListImgurl: "",
		},
		{
			ListId:     54,
			ListStatus: 0,
			ListText:   "一起睡个懒觉，赖个床🍧",
			ListImgurl: "",
		},
		{
			ListId:     55,
			ListStatus: 0,
			ListText:   "偷偷为对方买喜欢又舍不得的东西🍨",
			ListImgurl: "",
		},
		{
			ListId:     56,
			ListStatus: 0,
			ListText:   "一起坐一次巴士，在没去过的地方下车🍦",
			ListImgurl: "",
		},
		{
			ListId:     57,
			ListStatus: 0,
			ListText:   "一起为布置小家出主意🍰",
			ListImgurl: "",
		},
		{
			ListId:     58,
			ListStatus: 0,
			ListText:   "一起在午夜看一次恐怖片🎂",
			ListImgurl: "",
		},
		{
			ListId:     59,
			ListStatus: 0,
			ListText:   "一起去挑选一束花🧁",
			ListImgurl: "",
		},
		{
			ListId:     60,
			ListStatus: 0,
			ListText:   "一起去跳一次广场舞🥧",
			ListImgurl: "",
		},
		{
			ListId:     61,
			ListStatus: 0,
			ListText:   "一起为对方按摩一次🍮",
			ListImgurl: "",
		},
		{
			ListId:     62,
			ListStatus: 0,
			ListText:   "一起放一次风筝🍭",
			ListImgurl: "",
		},
		{
			ListId:     63,
			ListStatus: 0,
			ListText:   "一起吐槽一次对方的缺点🍬",
			ListImgurl: "",
		},
		{
			ListId:     64,
			ListStatus: 0,
			ListText:   "接对方下班一次🍫",
			ListImgurl: "",
		},
		{
			ListId:     65,
			ListStatus: 0,
			ListText:   "当陌生人一天，不许交流🍿",
			ListImgurl: "",
		},
		{
			ListId:     66,
			ListStatus: 0,
			ListText:   "为对方做便当🍩",
			ListImgurl: "",
		},
		{
			ListId:     67,
			ListStatus: 0,
			ListText:   "一起存钱🍪",
			ListImgurl: "",
		},
		{
			ListId:     68,
			ListStatus: 0,
			ListText:   "一起去看樱花🥮",
			ListImgurl: "",
		},
		{
			ListId:     69,
			ListStatus: 0,
			ListText:   "一起敷面膜🥠",
			ListImgurl: "",
		},
		{
			ListId:     70,
			ListStatus: 1,
			ListText:   "一起去一次动物园☕",
			ListImgurl: "https://img.gejiba.com/images/45090d8851732d1b90dbad90f1c4b887.jpg",
		},
		{
			ListId:     71,
			ListStatus: 1,
			ListText:   "一起骑行车🍵",
			ListImgurl: "https://img.gejiba.com/images/0e460ae2a5083e82b7c0051064443ec8.jpg",
		},
		{
			ListId:     72,
			ListStatus: 0,
			ListText:   "一起拍照洗照片贴房间🥣",
			ListImgurl: "",
		},
		{
			ListId:     73,
			ListStatus: 0,
			ListText:   "一起听一次演唱会🍼",
			ListImgurl: "",
		},
		{
			ListId:     74,
			ListStatus: 0,
			ListText:   "一起去一次酒吧🥤",
			ListImgurl: "",
		},
		{
			ListId:     75,
			ListStatus: 0,
			ListText:   "一起去听一次相声🧃",
			ListImgurl: "",
		},
		{
			ListId:     76,
			ListStatus: 0,
			ListText:   "一起玩一次真心话大冒险🧉",
			ListImgurl: "",
		},
		{
			ListId:     77,
			ListStatus: 0,
			ListText:   "一起去许愿池许个愿🥛",
			ListImgurl: "",
		},
		{
			ListId:     78,
			ListStatus: 0,
			ListText:   "一起入住一次五星级酒店，看夜景🍺",
			ListImgurl: "",
		},
		{
			ListId:     79,
			ListStatus: 0,
			ListText:   "一起去见父母🍻",
			ListImgurl: "",
		},
		{
			ListId:     80,
			ListStatus: 0,
			ListText:   "一起挑选戒指🍷",
			ListImgurl: "",
		},
		{
			ListId:     81,
			ListStatus: 0,
			ListText:   "一起挑选婚纱🥂",
			ListImgurl: "",
		},
		{
			ListId:     82,
			ListStatus: 0,
			ListText:   "一起为我们的小家添置东西🥃",
			ListImgurl: "",
		},
		{
			ListId:     83,
			ListStatus: 0,
			ListText:   "一起期待未来甜蜜小生活🍸",
			ListImgurl: "",
		},
	}
	return mgr.db.Create(&tl).RowsAffected
}

//关于
func (mgr *manager) CreateAbout() int64 {
	ab := model.About{
		Id:       1,
		Title:    "About",
		Aboutimg: "https://img.gejiba.com/images/7c5f1a655788c0d6ff0cae50232dde65.jpg",
		Info1:    "Hi, 欢迎你的来访",
		Info2:    "花有重开日 人无再少年",
		Info3:    "记录生活 留住感动",
		Btn1:     "听我介绍",
		Btn2:     "结束介绍",
		Infox1:   "情侣小站Like Girl for Golang 是我的二开重构项目",
		Infox2:   "在暑假假期最后的几天里写完了1.0版本",
		Infox3:   "目前已开源到码云 版本更新到4.0 项目迎来了尾声",
		Infox4:   "写后端的时候也是头都大了，不过幸好让我顺利写完了",
		Infox5:   "在开发过程中遇到了许多问题 也是只能自己探索解决 ...",
		Infox6:   "喜欢探索知识，热爱学习新知识，热爱开源文化",
		Btnx2:    "为什么叫 Ki？",
		Infof1:   "不知道你有没有看过《比悲伤更悲伤的故事》",
		Infof2:   "嗨，我是k，如果有下辈子的话，",
		Infof3:   "“我想当戒指，眼镜，床和笔记本，这样的话，我就可以...”",
		Infof4:   "当然跟这个没有关系哈哈",
		Btnf3:    "本站前端所有页面",
		Infod1:   "首页 index",
		Infod2:   "点点滴滴 little",
		Infod3:   "留言板 leaving",
		Infod4:   "关于 about",
		Infod5:   "欢迎您的来访 IP已记录",
	}
	return mgr.db.Create(&ab).RowsAffected
}

//用户
func (mgr *manager) CreateUser(id int, name string, pass string, qq string, sex int) int64 {
	user := model.User{
		Id:       id,
		Username: name,
		Password: pass,
		Qq:       qq,
		Sex:      sex,
	}
	return mgr.db.Create(&user).RowsAffected
}

//ip黑名单
func (mgr *manager) CreateIpBlackList() int64 {
	ip := model.IpBlackList{
		Id:      1,
		Ip:      "123.123.123",
		Time:    1718010394000,
		Commit:  "测试拦截",
		Address: "北京市",
	}
	return mgr.db.Create(&ip).RowsAffected
}
