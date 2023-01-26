package trending

import (
	"context"
	"net/http"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"

	"github.com/PuerkitoBio/goquery"
	"github.com/zeromicro/go-zero/core/logx"
)

type TrendingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTrendingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrendingLogic {
	return &TrendingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TrendingLogic) Trending(req *types.Trendingreque) (resp *types.Trendingrespo, err error) {
	var trendings []types.Trending
	var singletrending types.Trending
	titles,imgs,contents,hrefs:=l.GetNews("https://top.baidu.com/board?tab=realtime")
	for i:=0;i<len(titles);i++{
		singletrending.Tid=i
		singletrending.Tittle=titles[i]
		singletrending.Img=imgs[i]
		singletrending.Content=contents[i]
		singletrending.Href=hrefs[i]
		trendings=append(trendings, singletrending)
	}
	return &types.Trendingrespo{
		Error_code: 0,
		Trending: trendings,
	},nil
}
//The returns are :
//Tittles  Imgs   Contents Hrefs
func(l *TrendingLogic)GetNews(url string)([]string,[]string,[]string,[]string){
	var Tittle,Img,Content,Href []string
	client:=&http.Client{}
	requ,_:=http.NewRequest("GET",url,nil)
	respo,_:=client.Do(requ)
	docu,_:=goquery.NewDocumentFromReader(respo.Body)
	docu.Find("#sanRoot > main > div.container.right-container_2EFJr > div > div:nth-child(2) > div > div.content_1YWBm > a > div.c-single-text-ellipsis").Each(func (i int,s *goquery.Selection)  {
		Tittle=append(Tittle, s.Text())
	})
	docu.Find("#sanRoot > main > div.container.right-container_2EFJr > div > div:nth-child(2) > div > a > img").Each(func(i int, s *goquery.Selection) {
		img,_:=s.Attr("src")
		Img=append(Img, img)
	})
	docu.Find("#sanRoot > main > div.container.right-container_2EFJr > div > div:nth-child(2) > div > div.content_1YWBm > div.hot-desc_1m_jR.small_Uvkd3").Each(func(i int, s *goquery.Selection) {
		href,_:=s.Find("a").Attr("href")
		content:=s.Text()
		Href=append(Href, href)
		Content=append(Content, content)
	})
	return Tittle,Img,Content,Href
}
