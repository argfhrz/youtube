from requests_html import HTMLSession
from bs4 import BeautifulSoup as bs


# init session
session = HTMLSession()


def get_video_info():

    video = ["https://www.youtube.com/watch?v=3o0t5Sy_3yw",
    "https://www.youtube.com/watch?v=JQxQqdn9dEM",
    "https://www.youtube.com/watch?v=ppjO2aYR4kQ",
    "https://www.youtube.com/watch?v=X2qDtU7oyPw",
    "https://www.youtube.com/watch?v=aOtHnJcZiLo",
    "https://www.youtube.com/watch?v=dkL15Dq9Rho",
    "https://www.youtube.com/watch?v=tg2uF3R_Ozo",
    "https://www.youtube.com/watch?v=SPvxgkziGds",
    "https://www.youtube.com/watch?v=r8T61-29o10",
    "https://www.youtube.com/watch?v=ZBXqydJ-kD0",
    "https://www.youtube.com/watch?v=yW86RTjyDE0",
    "https://www.youtube.com/watch?v=Psr7yfO_WY0",
    "https://www.youtube.com/watch?v=3Lfs0mFkxQE",
    "https://www.youtube.com/watch?v=qnZN5oY3KmY",
    "https://www.youtube.com/watch?v=bWj3M9QU5Mo",
    "https://www.youtube.com/watch?v=OrTvH8a7bkQ",
    "https://www.youtube.com/watch?v=s_hk-nszOmg",
    "https://www.youtube.com/watch?v=k5DRcpOYbB4",
    "https://www.youtube.com/watch?v=QuWple6WhE4",
    "https://www.youtube.com/watch?v=7dP_PmfeHQY"]
    result = {"channelID":[], "title":[], "channelName":[], "datePublished":[]}
    for x in video:
        
        response = session.get(x)
        
        response.html.render(timeout=60)
       
        soup = bs(response.html.html, "html.parser")
      
        arg = soup.find("yt-formatted-string", {"class": "ytd-channel-name"}).find("a")
        arg = arg.text
       
        #channel ID
        channelID = soup.find("meta", itemprop="channelId")['content']
        result["channelID"].append(channelID)
        # video title
        title = soup.find("meta", itemprop="name")['content'] 
        result["title"].append(title)
        # channel details
        result["channelName"].append(arg)
        #publish date
        date = soup.find("meta", itemprop="datePublished")['content']
        result["datePublished"].append(date)
        

    return result
        


if __name__ == "__main__":
    data = get_video_info()
    




  