import json
import datetime


def compare_json(json_upcoming, json_str2):
    # Convert JSON strings to Python dictionaries
    # dict1 = json.loads(json_str1)
    dict2 = json.loads(json_str2)

    # Sort keys for each dictionary
    dict1_sorted = dict(sorted(json_upcoming.items()))
    dict2_sorted = dict(sorted(dict2.items()))

    # Compare dictionaries
    print("\ndict1_sorted == dict2_sorted", dict1_sorted == dict2_sorted)
    print("json_upcoming", dict1_sorted)
    print("json_existing", dict2_sorted)
    return dict1_sorted == dict2_sorted


json_upcoming = {"barCode": "4897041441949", "brand": "Smile Makers", "catalog_name": "Health and Beauty@Wellness", "currency": "EUR", "gender": "female", "imageUrls": [{"class1": "详情", "class2": 1, "class3": "", "url": "https://s4.thcdn.com//productimg/1600/1600/14866183-7955070894878451.jpg"}, {"class1": "详情", "class2": 2, "class3": "", "url": "https://s4.thcdn.com//productimg/1600/1600/14866183-9535070894930123.jpg"}, {"class1": "详情", "class2": 3, "class3": "", "url": "https://s4.thcdn.com//productimg/1600/1600/14866183-7335070894992031.jpg"}, {"class1": "详情", "class2": 4, "class3": "", "url": "https://s4.thcdn.com//productimg/1600/1600/14866183-1235070895041726.jpg"}], "introduction": [
    {"title": "description", "content": "Hit the spot with the Smile Makers The Tennis Pro G-Spot Vibrator. Ideal for couples or newcomers, this multitasking toy is a g-spot massager and clitoral vibrator combined, allowing you to experience a breathtaking blended orgasm with every serve.  Made from a smooth skin-like silicone, it features an angled and rounded head with four speeds and two pulsation modes, ideal for internal and external stimulation. The battery-powered tool is small but mighty and doesn’t make a ‘racket’."}], "originalPrice": 57.95, "price": 57.95, "productUrl": "https://www.lookfantastic.fr/smile-makers-the-tennis-pro-g-spot-vibrator/14866183.html?switchcurrency=EUR&shippingcountry=FR", "product_name": "Smile Makers The Tennis Pro G-Spot Vibrator", "productcode": "/smile-makers-the-tennis-pro-g-spot-vibrator/14866183.html", "promotionTitle": "", "reference": "14866183", "series": "14866183", "shopCode": "LOOKFANTASTIC FR", "site_id": 7, "specs": [], "stock": 1}
json_existing = '{"barCode": "4897041441949", "brand": "Smile Makers", "catalog_name": "Health and Beauty@Wellness", "currency": "EUR", "gender": "female", "imageUrls": [{"url": "https://s4.thcdn.com//productimg/1600/1600/14866183-7955070894878451.jpg", "class1": "详情", "class2": 1, "class3": ""}, {"url": "https://s4.thcdn.com//productimg/1600/1600/14866183-9535070894930123.jpg", "class1": "详情", "class2": 2, "class3": ""}, {"url": "https://s4.thcdn.com//productimg/1600/1600/14866183-7335070894992031.jpg", "class1": "详情", "class2": 3, "class3": ""}, {"url": "https://s4.thcdn.com//productimg/1600/1600/14866183-1235070895041726.jpg", "class1": "详情", "class2": 4, "class3": ""}], "introduction": [{"title": "description", "content": "Hit the spot with the Smile Makers The Tennis Pro G-Spot Vibrator. Ideal for couples or newcomers, this multitasking toy is a g-spot massager and clitoral vibrator combined, allowing you to experience a breathtaking blended orgasm with every serve.  Made from a smooth skin-like silicone, it features an angled and rounded head with four speeds and two pulsation modes, ideal for internal and external stimulation. The battery-powered tool is small but mighty and doesn’t make a ‘racket’."}], "originalPrice": 57.95, "price": 57.95, "productUrl": "https://www.lookfantastic.fr/smile-makers-the-tennis-pro-g-spot-vibrator/14866183.html?switchcurrency=EUR&shippingcountry=FR", "product_name": "Smile Makers The Tennis Pro G-Spot Vibrator", "productcode": "/smile-makers-the-tennis-pro-g-spot-vibrator/14866183.html", "promotionTitle": "", "reference": "14866183", "series": "14866183", "shopCode": "LOOKFANTASTIC FR", "site_id": 7, "specs": [], "stock": 1}'
print(type(json_upcoming), type(json_existing))
print(compare_json(json_upcoming, json_existing))

if compare_json(json_upcoming, json_existing):
    print("false", datetime.datetime.now())
else:
    print("true")
