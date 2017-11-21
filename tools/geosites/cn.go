package geosites

import (
	"v2ray.com/core/app/router"
)

var (
	chinaSitesDomains []*router.Domain
)

func GetGeoSiteCN() []*router.Domain {
	return chinaSitesDomains
}

func init() {
	domains := []string{
		"cn",
		"xn--fiqs8s", /* .中国 */

		"10010.com",
		"100offer.com",
		"115.com",
		"123juzi.com",
		"123juzi.net",
		"123u.com",
		"126.com",
		"126.net",
		"127.net",
		"163.com",
		"17173.com",
		"17cdn.com",
		"188.com",
		"1905.com",
		"21cn.com",
		"2288.org",
		"2345.com",
		"263.net",
		"2cto.com",
		"3322.org",
		"35.com",
		"360doc.com",
		"360buy.com",
		"360buyimg.com",
		"360safe.com",
		"36kr.com",
		"39.net",
		"3dmgame.com",
		"3conline.com",
		"4399.com",
		"500d.me",
		"50bang.org",
		"51.la",
		"51credit.com",
		"51cto.com",
		"51jingying.com",
		"51job.com",
		"51jobcdn.com",
		"51wendang.com",
		"55.com",
		"51yes.com",
		"55bbs.com",
		"58.com",
		"6rooms.com",
		"71.am",
		"7k7k.com",
		"900.la",
		"9718.com",
		"9xu.com",
		"abchina.com",
		"acfun.tv",
		"acgvideo.com",
		"agrantsem.com",
		"aicdn.com",
		"aixifan.com",
		"alibaba.com",
		"alibabaplanet.com",
		"alicdn.com",
		"aliimg.com.com",
		"alipay.com",
		"alipayobjects.com",
		"aliyun.com",
		"aliyuncdn.com",
		"aliyuncs.com",
		"allyes.com",
		"amap.com",
		"anjuke.com",
		"anquan.org",
		"appinn.com",
		"babytree.com",
		"babytreeimg.com",
		"baidu.com",
		"baiducontent.com",
		"baidupcs.com",
		"baidustatic.com",
		"baifendian.com",
		"baifubao.com",
		"baihe.com",
		"baike.com",
		"baixing.com",
		"baixing.net",
		"bankcomm.com",
		"bankofchina.com",
		"bcy.net",
		"bdimg.com",
		"bdstatic.com",
		"bilibili.com",
		"cn.bing.com",
		"bitauto.com",
		"bitautoimg.com",
		"bobo.com",
		"bootcss.com",
		"btcfans.com",
		"caiyunapp.com",
		"ccb.com",
		"cctv.com",
		"cctvpic.com",
		"cdn20.com",
		"cebbank.com",
		"ch.com",
		"chashebao.com",
		"che168.com",
		"china.com",
		"chinacache.com",
		"chinacache.net",
		"chinahr.com",
		"chinamobile.com",
		"chinapay.com",
		"chinatranslation.net",
		"chinaz.com",
		"chiphell.com",
		"chouti.com",
		"chuangxin.com",
		"chuansong.me",
		"clouddn.com",
		"cloudxns.com",
		"cmbchina.com",
		"cnbeta.com",
		"cnbetacdn.com",
		"cnblogs.com",
		"cnepub.com",
		"cnzz.com",
		"coding.net",
		"coolapk.com",
		"cqvip.com",
		"csbew.com",
		"csdn.net",
		"ctfile.com",
		"ctrip.com",
		"cubead.com",
		"dajie.com",
		"dajieimg.com",
		"dangdang.com",
		"daocloud.io",
		"daovoice.io",
		"dbank.com",
		"dedecms.com",
		"dgtle.com",
		"diandian.com",
		"dianping.com",
		"diopic.net",
		"docin.com",
		"dockerone.com",
		"dockone.io",
		"donews.com",
		"douban.com",
		"douban.fm",
		"doubanio.com",
		"dpfile.com",
		"duokanbox.com",
		"duomai.com",
		"duoshuo.com",
		"duowan.com",
		"dxpmedia.com",
		"eastday.com",
		"ecitic.com",
		"emarbox.com",
		"eoeandroid.com",
		"etao.com",
		"excelhome.net",
		"fanli.com",
		"feng.com",
		"fengniao.com",
		"fhldns.com",
		"foxmail.com",
		"geekpark.net",
		"geetest.com",
		"geilicdn.com",
		"getui.com",
		"google-analytics.com",
		"growingio.com",
		"gtags.net",
		"gwdang.com",
		"hao123.com",
		"hao123img.com",
		"haosou.com",
		"hdslb.com",
		"henha.com",
		"henkuai.com",
		"hexun.com",
		"hichina.com",
		"huanqiu.com",
		"hunantv.com",
		"huochepiao.com",
		"hupu.com",
		"hupucdn.com",
		"huxiu.com",
		"iask.com",
		"iciba.com",
		"idqqimg.com",
		"ifanr.com",
		"ifanrusercontent.com",
		"ifanrx.com",
		"ifeng.com",
		"ifengimg.com",
		"ijinshan.com",
		"ikafan.com",
		"imedao.com",
		"imgo.tv",
		"imooc.com",
		"infoq.com",
		"infoqstatic.com",
		"ip138.com",
		"ipinyou.com",
		"ipip.net",
		"ip-cdn.com",
		"iqiyi.com",
		"irs01.com",
		"it165.net",
		"it168.com",
		"it610.com",
		"iteye.com",
		"ithome.com",
		"itjuzi.com",
		"jandan.net",
		"jd.com",
		"jb51.com",
		"jia.com",
		"jianshu.com",
		"jianshu.io",
		"jiasuhui.com",
		"jiathis.com",
		"jiayuan.com",
		"jikexueyuan.com",
		"jisuanke.com",
		"jmstatic.com",
		"jsdelivr.net",
		"jstv.com",
		"jumei.com",
		"jyimg.com",
		"kaixin001.com",
		"kanimg.com",
		"kankanews.com",
		"kejet.net",
		"kf5.com",
		"kimiss.com",
		"kouclo.com",
		"koudai.com",
		"koudai8.com",
		"ku6.com",
		"ku6cdn.com",
		"ku6img.com",
		"kuqin.com",
		"lady8844.com",
		"lagou.com",
		"le.com",
		"leanote.com",
		"leiphone.com",
		"leju.com",
		"leturich.org",
		"letv.com",
		"letvcdn.com",
		"letvimg.com",
		"liantu.me",
		"liaoxuefeng.com",
		"liba.com",
		"libaclub.com",
		"liepin.com",
		"lietou.com",
		"lightonus.com",
		"linkvans.com",
		"linuxidc.com",
		"liuxiaoer.com",
		"lofter.com",
		"lu.com",
		"lufax.com",
		"lufaxcdn.com",
		"lvmama.com",
		"lxdns.com",
		"lxway.com",
		"ly.com",
		"mayihr.com",
		"mechina.org",
		"mediav.com",
		"meiqia.com",
		"meika360.com",
		"meilishuo.com",
		"meishij.net",
		"meituan.com",
		"meizu.com",
		"mgtv.com",
		"mi.com",
		"miaopai.com",
		"miaozhen.com",
		"miui.com",
		"mmbang.com",
		"mmbang.info",
		"mmstat.com",
		"mogucdn.com",
		"mogujie.com",
		"mop.com",
		"mscbsc.com",
		"mukewang.com",
		"mydrivers.com",
		"myshow360.net",
		"mzstatic.com",
		"netease.com",
		"newbandeng.com",
		"ngacn.cc",
		"ntalker.com",
		"nvsheng.com",
		"oeeee.com",
		"ol-img.com",
		"oneapm.com",
		"onlinedown.net",
		"onlinesjtu.com",
		"oschina.net",
		"paipai.com",
		"pcbeta.com",
		"pchome.net",
		"pingan.com",
		"pingplusplus.com",
		"pps.tv",
		"psbc.com",
		"pubyun.com",
		"qbox.me",
		"qcloud.com",
		"qhimg.com",
		"qhres.com",
		"qiaobutang.com",
		"qidian.com",
		"qie.tv",
		"qihucdn.com",
		"qingcloud.com",
		"qingsongchou.com",
		"qiniu.com",
		"qiniucdn.com",
		"qiniudn.com",
		"qiniudns.com",
		"qiyi.com",
		"qiyipic.com",
		"qtmojo.com",
		"qq.com",
		"qqmail.com",
		"qunar.com",
		"qunarzz.com",
		"qzone.com",
		"renren.com",
		"runoob.com",
		"ruanmei.com",
		"ruby-china.org",
		"sandai.net",
		"sanguosha.com",
		"sanwen.net",
		"segmentfault.com",
		"sf-express.com",
		"sharejs.com",
		"shmetro.com",
		"shutcm.com",
		"simei8.com",
		"sina.com",
		"sinaapp.com",
		"sinaedge.com",
		"sinaimg.com",
		"sinajs.com",
		"szzfgjj.com",
		"smzdm.com",
		"sohu.com",
		"sohucs.com",
		"sogou.com",
		"sogoucdn.com",
		"soso.com",
		"sspai.com",
		"ssports.com",
		"starbaby.cc",
		"starbaby.com",
		"staticfile.org",
		"stockstar.com",
		"suning.com",
		"szfw.org",
		"t1y5.com",
		"tanx.com",
		"tao123.com",
		"taobao.com",
		"taobaocdn.com",
		"tbcache.com",
		"tencent.com",
		"tenpay.com",
		"tenxcloud.com",
		"tiebaimg.com",
		"tietuku.com",
		"tiexue.net",
		"tmall.com",
		"tmcdn.net",
		"topthink.com",
		"ttpod.com",
		"tudou.com",
		"tudouui.com",
		"tuicool.com",
		"tuniu.com",
		"tutuapp.com",
		"u17.com",
		"useso.com",
		"unionpay.com",
		"unionpaysecure.com",
		"upyun.com",
		"upaiyun.com",
		"v2ex.com",
		"v5875.com",
		"vamaker.com",
		"vancl.com",
		"vcimg.com",
		"vip.com",
		"wallstreetcn.com",
		"wandoujia.com",
		"wdjimg.com",
		"weand.com",
		"webterren.com",
		"weibo.com",
		"weicaifu.com",
		"weidian.com",
		"weiphone.com",
		"weiphone.net",
		"weixing.com",
		"weiyun.com",
		"wonnder.com",
		"worktile.com",
		"wooyun.org",
		"wrating.com",
		"wscdns.com",
		"wumii.com",
		"xiachufang.com",
		"xiami.com",
		"xiaokaxiu.com",
		"xiaomi.com",
		"xitu.com",
		"xinhuanet.com",
		"xinshipu.com",
		"xiu8.com",
		"xnpic.com",
		"xueqiu.com",
		"xunlei.com",
		"xywy.com",
		"yaolan.com",
		"yccdn.com",
		"yeepay.com",
		"yesky.com",
		"yigao.com",
		"yihaodian.com",
		"yihaodianimg.com",
		"yingjiesheng.com",
		"yinxiang.com",
		"yinyuetai.com",
		"yixi.tv",
		"yjbys.com",
		"yhd.com",
		"youboy.com",
		"youku.com",
		"ysten.com",
		"yunba.io",
		"yundaex.com",
		"yunshipei.com",
		"yupoo.com",
		"yuzua.com",
		"yy.com",
		"yytcdn.com",
		"zampda.net",
		"zastatic.com",
		"zbjimg.com",
		"zdfans.com",
		"zdmimg.com",
		"zhenai.com",
		"zhanqi.tv",
		"zhaopin.com",
		"zhihu.com",
		"zhimg.com",
		"zhiziyun.com",
		"zjstv.com",
		"zhubajie.com",
		"zrblog.net",
		"zuche.com",
		"zuchecdn.com",
	}

	chinaSitesDomains = make([]*router.Domain, len(domains))
	for idx, pattern := range domains {
		chinaSitesDomains[idx] = &router.Domain{
			Type:  router.Domain_Domain,
			Value: pattern,
		}
	}
}
