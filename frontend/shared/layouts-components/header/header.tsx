
import Link from "next/link";
import React, { Fragment, useEffect, useRef, useState } from "react";
import { ThemeChanger, removeFromCart } from "@/shared/redux/action";
import store from "@/shared/redux/store";
import { connect, useDispatch, useSelector } from "react-redux";
import { MENUITEMS } from "../sidebar/nav";
import SimpleBar from "simplebar-react";
import SpkBadge from "@/shared/@spk-reusable-components/uielements/spk-badge";
import SpkDropdown from "@/shared/@spk-reusable-components/uielements/spk-dropdown";
import nextConfig from "@/next.config";
import SpkButton from "@/shared/@spk-reusable-components/uielements/spk-button";
import { useLanguage } from "@/shared/i18n/LanguageContext";

const Header = ({ local_varaiable, ThemeChanger }: any) => {

  let { basePath } = nextConfig;

  const { lang, setLanguage } = useLanguage();

  //full screen
  const [fullScreen, setFullScreen] = useState(false);

  const toggleFullScreen = () => {
    const elem = document.documentElement;

    if (!document.fullscreenElement) {
      elem.requestFullscreen().then(() => setFullScreen(true));
    } else {
      document.exitFullscreen().then(() => setFullScreen(false));
    }
  };

  //MenuClose

  function menuClose() {
    const theme = store.getState();
    if (window.innerWidth <= 992) {
      ThemeChanger({ ...theme, toggled: "close" });
    }
    if (window.innerWidth >= 992) {
      ThemeChanger({
        ...theme,
        toggled: local_varaiable.toggled ? local_varaiable.toggled : "",
      });
    }
  }

  //Toggle-Function

  const overlayRef = useRef<HTMLDivElement | null>(null);

  function toggleSidebar() {
    const theme = store.getState();
    let sidemenuType = theme.dataNavLayout;
    if (window.innerWidth >= 992) {
      if (sidemenuType === "vertical") {
        let verticalStyle = theme.dataVerticalStyle;
        const navStyle = theme.dataNavStyle;
        switch (verticalStyle) {
          // closed
          case "closed":
            ThemeChanger({ ...theme, dataNavStyle: "" });
            if (theme.toggled === "close-menu-close") {
              ThemeChanger({ ...theme, toggled: "" });
            } else {
              ThemeChanger({ ...theme, toggled: "close-menu-close" });
            }
            break;
          // icon-overlay
          case "overlay":
            ThemeChanger({ ...theme, dataNavStyle: "" });
            if (theme.toggled === "icon-overlay-close") {
              ThemeChanger({ ...theme, toggled: "", iconOverlay: "" });
            } else {
              if (window.innerWidth >= 992) {
                ThemeChanger({
                  ...theme,
                  toggled: "icon-overlay-close",
                  iconOverlay: "",
                });
              }
            }
            break;
          // icon-text
          case "icontext":
            ThemeChanger({ ...theme, dataNavStyle: "" });
            if (theme.toggled === "icon-text-close") {
              ThemeChanger({ ...theme, toggled: "" });
            } else {
              ThemeChanger({ ...theme, toggled: "icon-text-close" });
            }
            break;
          // doublemenu
          case "doublemenu":
            ThemeChanger({ ...theme, dataNavStyle: "" });
            ThemeChanger({ ...theme, dataNavStyle: "" });
            if (theme.toggled === "double-menu-open") {
              ThemeChanger({ ...theme, toggled: "double-menu-close" });
            } else {
              let sidemenu = document.querySelector(".side-menu__item.active");
              if (sidemenu) {
                ThemeChanger({ ...theme, toggled: "double-menu-open" });
                if (sidemenu.nextElementSibling) {
                  sidemenu.nextElementSibling.classList.add(
                    "double-menu-active"
                  );
                } else {
                  ThemeChanger({ ...theme, toggled: "double-menu-close" });
                }
              }
            }
            break;
          // detached
          case "detached":
            if (theme.toggled === "detached-close") {
              ThemeChanger({ ...theme, toggled: "", iconOverlay: "" });
            } else {
              ThemeChanger({
                ...theme,
                toggled: "detached-close",
                iconOverlay: "",
              });
            }

            break;

          // default
          case "default":
            ThemeChanger({ ...theme, toggled: "" });
        }
        switch (navStyle) {
          case "menu-click":
            if (theme.toggled === "menu-click-closed") {
              ThemeChanger({ ...theme, toggled: "" });
            } else {
              ThemeChanger({ ...theme, toggled: "menu-click-closed" });
            }
            break;
          // icon-overlay
          case "menu-hover":
            if (theme.toggled === "menu-hover-closed") {
              ThemeChanger({ ...theme, toggled: "" });
            } else {
              ThemeChanger({ ...theme, toggled: "menu-hover-closed" });
            }
            break;
          case "icon-click":
            if (theme.toggled === "icon-click-closed") {
              ThemeChanger({ ...theme, toggled: "" });
            } else {
              ThemeChanger({ ...theme, toggled: "icon-click-closed" });
            }
            break;
          case "icon-hover":
            if (theme.toggled === "icon-hover-closed") {
              ThemeChanger({ ...theme, toggled: "" });
            } else {
              ThemeChanger({ ...theme, toggled: "icon-hover-closed" });
            }
            break;
        }
      }
    } else {
      if (theme.toggled === "close") {
        ThemeChanger({ ...theme, toggled: "open" });



        setTimeout(() => {
          if (theme.toggled === "open") {
            if (overlayRef.current) {
              overlayRef.current.classList.remove("active");
            }

            if (overlayRef.current) {
              overlayRef.current.classList.add("active");
              overlayRef.current.addEventListener("click", () => {
                if (overlayRef.current) {
                  overlayRef.current.classList.remove("active");
                  menuClose();
                }
              });
            }
          }

          const resizeHandler = () => {
            if (window.innerWidth >= 992) {
              if (overlayRef.current) {
                overlayRef.current.classList.remove("active");
              }
            }
          };

          window.addEventListener("resize", resizeHandler);

          // Cleanup function to remove event listeners
          return () => {
            window.removeEventListener("resize", resizeHandler);
          };

        }, 100);
      } else {
        ThemeChanger({ ...theme, toggled: "close" });
      }
    }
  }

  //Dark Model

  const toggledark = () => {

    ThemeChanger({
      ...local_varaiable,
      "class": local_varaiable.class == 'dark' ? 'light' : 'dark',
      "dataHeaderStyles": local_varaiable.class == 'dark' ? 'light' : 'dark',
      "dataMenuStyles": local_varaiable.dataNavLayout == 'horizontal' ? local_varaiable.class == 'dark' ? 'light' : 'dark' : "dark"

    });
    const theme = store.getState();

    if (theme.class != 'dark') {

      ThemeChanger({
        ...theme,
        "bodyBg": '',
        "lightRgb": '',
        "darkBg": '',
        "inputBorder": '',
        "gray": '',
      });
      localStorage.setItem("xintralighttheme", "light");
      localStorage.removeItem("xintradarktheme");
      localStorage.removeItem("bodyBg");
      localStorage.removeItem("darkBg");
      localStorage.removeItem("xintraMenu");
      localStorage.removeItem("xintraHeader");
      localStorage.removeItem("hs_theme");
    }
    else {
      localStorage.setItem("xintradarktheme", "dark");
      localStorage.removeItem("xintralighttheme");
      localStorage.removeItem("xintraMenu");
      localStorage.removeItem("xintraHeader");
      localStorage.removeItem("hs_theme");
    }

  };

  //Switcher-Icon




  //Responsive Search
  //Notification
  const img = <span className="avatar avatar-md avatar-rounded bg-primary"> <img src={`${process.env.NODE_ENV === 'production' ? basePath : ''}/assets/images/faces/1.jpg`} alt="user1" /> </span>
  const img1 = <span className="avatar avatar-md bg-primary avatar-rounded fs-20"> <i className="fe fe-shopping-cart leading-none"></i> </span>
  const img2 = <span className="avatar avatar-md bg-orange avatar-rounded"> <img src={`${process.env.NODE_ENV === 'production' ? basePath : ''}/assets/images/faces/10.jpg`} alt="user1" /> </span>
  const img3 = <span className="avatar avatar-md bg-success avatar-rounded"> <img src={`${process.env.NODE_ENV === 'production' ? basePath : ''}/assets/images/faces/11.jpg`} alt="user1" /> </span>
  const img4 = <span className="avatar avatar-md bg-primarytint2color avatar-rounded"> <i className="ri-gift-line leading-none text-[1rem]"></i> </span>

  const content = <div className="text-muted fw-normal fs-12 header-notification-text text-truncate">Jane Sam sent you a message.</div>
  const content1 = <div className="text-muted fw-normal fs-12 header-notification-text text-truncate">Order <span className="text-primarytint1color">#54321</span> has been shipped.</div>
  const content2 = <div className="text-muted fw-normal fs-12 header-notification-text text-truncate">Reacted: <span className="text-primarytint3color">John Richard</span> on your next purchase!</div>
  const content3 = <div className="text-muted fw-normal fs-12 header-notification-text text-truncate"><span className="text-info">Kelin Brown</span> has sent you the request.</div>
  const content4 = <div className="text-muted fw-normal fs-12 header-notification-text text-truncate">Enjoy<span className="text-success">   20% off</span> on your next purchase!</div>

  const Notifications = [
    { id: 1, src: img, heading: "New Messages", data: content, data1: "Now" , },
    { id: 2, src: img1, heading: "Order Updates", data: content1, data1: "2 hours ago" , },
    { id: 3, src: img2, heading: "Comment on Post", data: content2, data1: "2 hours ago" ,},
    { id: 4, src: img3, heading: "Follow Request", data: content3, data1: "1 Day ago" , },
    { id: 5, src: img4, heading: "Exclusive Offers", data: content4, data1: "5 hours ago" , },
  ]

  const [notifications, setNotifications] = useState(Notifications); // assuming 'data' is an array of notifications
  const [unreadCount, setUnreadCount] = useState(5); // initial unread count

  const hasNotifications = notifications.length > 0;

  const handleRemove1 = (id: number) => {
    if (event) {
      event.stopPropagation();
    }
    // Filter out the notification by id
    const updatedNotifications = notifications.filter((notification) => notification.id !== id);
    setNotifications(updatedNotifications);
    setUnreadCount(unreadCount - 1); // decrease unread count
  };

  //Cart function

  const cartProduct = [
    {
      id: 1,
      productpicture: "/assets/images/ecommerce/png/30.png",
      title: "Wireless Headphones",
      price: "$75",
      discount: "$75",
      quantity: "1",
      data: "Wireless freedom with crystal-clear sound and comfortable "
    },
    {
      id: 2,
      productpicture: "/assets/images/ecommerce/png/29.png",
      title: "Ladies Hand Bag",
      price: "$15",
      discount: "$30",
      quantity: "2",
      data: "Both fashion and functionality. "
    },
    {
      id: 3,
      productpicture: "/assets/images/ecommerce/png/32.png",
      title: "Alarm Clock",
      price: "$84",
      discount: "$84",
      quantity: "1",
      data: "Add natural beauty to your space "
    },
    {
      id: 4,
      productpicture: "/assets/images/ecommerce/png/12.png",
      title: "Kids' Party Wear Frock",
      price: "$37",
      discount: "$37",
      quantity: "1",
      data: "Crafted from soft, breathable fabric and adorned with delightful "
    },
    {
      id: 5,
      productpicture: "/assets/images/ecommerce/png/16.png",
      title: "Advanced Smart Watch",
      price: "$29",
      discount: "$48",
      quantity: "2",
      data: "ultimate in wearable technology,combining cutting-edge "
    },
  ];

  const maxDisplayItems = 5;

  const dispatch = useDispatch();
  const reduxCart = useSelector((state: any) => state.cart);
  const [localCart, setLocalCart] = useState(cartProduct);
  const [remainingCount2, setRemainingCount2] = useState(0);

  // Combine local and redux carts whenever they change
  const card = [...localCart, ...reduxCart];

  useEffect(() => {
    setRemainingCount2(card.length);
    setCartItemCount(localCart.length);
  }, [card, localCart]);

  const handleDelete = (id: number, event: any) => {
    event.stopPropagation();
    // Remove item from local cart
    const updatedLocalCart = localCart.filter(item => item.id !== id);
    setLocalCart(updatedLocalCart);
    // Update cart item count
    setCartItemCount(updatedLocalCart.length);
    // Remove item from redux cart
    dispatch(removeFromCart(id));
  };
  const [cartItems, setCartItems] = useState([...cartProduct]);
  const [cartItemCount, setCartItemCount] = useState(cartProduct.length);

  useEffect(() => {
    setCartItemCount(localCart.length);
  }, [localCart]);

  const handleRemove = (itemId: number, event: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => {
    event.stopPropagation();
    const updatedCart = localCart.filter((item) => item.id !== itemId);
    setLocalCart(updatedCart);
    setCartItemCount(updatedCart.length);
  };


  return (
    <Fragment>
      <header className="app-header" id="header">
        {/* <!-- Start::main-header-container --> */}
        <div className="main-header-container container-fluid">
          {/* <!-- Start::header-content-left --> */}
          <div className="header-content-left">
            {/* <!-- Start::header-element --> */}
            <div className="header-element">
              <div className="horizontal-logo">
                <Link href="/components/dashboard/sales" className="header-logo">
                  <img
                    src={`${process.env.NODE_ENV === "production" ? basePath : ""
                      }/assets/images/brand-logos/desktop-logo.png`}
                    alt="logo"
                    className="desktop-logo"
                  />
                  <img
                    src={`${process.env.NODE_ENV === "production" ? basePath : ""
                      }/assets/images/brand-logos/toggle-dark.png`}
                    alt="logo"
                    className="toggle-dark"
                  />
                  <img
                    src={`${process.env.NODE_ENV === "production" ? basePath : ""
                      }/assets/images/brand-logos/desktop-dark.png`}
                    alt="logo"
                    className="desktop-dark"
                  />
                  <img
                    src={`${process.env.NODE_ENV === "production" ? basePath : ""
                      }/assets/images/brand-logos/toggle-logo.png`}
                    alt="logo"
                    className="toggle-logo"
                  />
                  <img
                    src={`${process.env.NODE_ENV === "production" ? basePath : ""
                      }/assets/images/brand-logos/toggle-white.png`}
                    alt="logo"
                    className="toggle-white"
                  />
                  <img
                    src={`${process.env.NODE_ENV === "production" ? basePath : ""
                      }/assets/images/brand-logos/desktop-white.png`}
                    alt="logo"
                    className="desktop-white"
                  />
                </Link>
              </div>
            </div>
            {/* <!-- End::header-element --> */}

            {/* <!-- Start::header-element --> */}
            <div className="header-element mx-lg-0">
              <Link
                aria-label="Hide Sidebar"
                onClick={() => toggleSidebar()}
                className="sidemenu-toggle header-link animated-arrow hor-toggle horizontal-navtoggle"
                data-bs-toggle="sidebar"
                scroll={false}
                href="#!"
              >
                <span></span>
              </Link>
            </div>
            {/* <!-- End::header-element --> */}

            {/* <!-- Start::header-element --> */}

            <div  className="header-element header-search md:!block !hidden my-auto auto-complete-search" id="search-modal">
              {/* Start::header-link */}
              <input
                type="text"
                className="header-search-bar form-control"
                id="header-search"
                placeholder="Search for Results..."
                autoComplete="off"
                autoCapitalize="off"
              />
              <Link
                scroll={false}
                href="#!"
                className="header-search-icon border-0"
              >
                <i className="ri-search-line"></i>
              </Link>
            </div>
            {/* <!-- End::header-element --> */}
          </div>
          {/* <!-- End::header-content-left --> */}

          {/* <!-- Start::header-content-right --> */}
          <ul className="header-content-right">
            {/* <!-- Start::header-element --> */}
            <li className="header-element md:!hidden block">
              <a
                aria-label="anchor"
                href="#!"
                className="header-link"
                data-bs-toggle="modal"
                data-hs-overlay="#header-responsive-search"
              >
                {/* <!-- Start::header-link-icon --> */}
                <i className="bi bi-search header-link-icon"></i>
                {/* <!-- End::header-link-icon --> */}
              </a>
            </li>
            {/* <!-- End::header-element --> */}

            {/* <!-- Start::header-element --> */}
            <SpkDropdown Customclass="header-element country-selector hidden sm:block [--placement:bottom-right] rtl:[--placement:bottom-left]" Linktag={true} Navigate='#!' Customtoggleclass='header-link hs-dropdown-toggle ti-dropdown-toggle' Svgicon='m10.5 21 5.25-11.25L21 21m-9-3h7.5M3 5.621a48.474 48.474 0 0 1 6-.371m0 0c1.12 0 2.233.038 3.334.114M9 5.25V3m3.334 2.364C11.176 10.658 7.69 15.08 3 17.502m9.334-12.138c.896.061 1.785.147 2.666.257m-4.589 8.495a18.023 18.023 0 0 1-3.827-5.802'
              SvgClass='w-6 h-6 header-link-icon' Linkclass='header-link hs-dropdown-toggle ti-dropdown-toggle' Svg={true} Custommenuclass='main-header-dropdown min-w-[10rem] hidden' SvgStroke="currentColor" Strokewidth="1.5" Svvgviewbox="0 0 24 24">
              <li>
                <a
                  className={`ti-dropdown-item flex items-center ${lang === "es" ? "bg-primary/10" : ""}`}
                  href="#!"
                  onClick={(e: React.MouseEvent) => { e.preventDefault(); setLanguage("es"); }}
                >
                  <div className="flex items-center">
                    <span className="avatar avatar-rounded avatar-xs leading-none me-2">
                      <img
                        src={`${process.env.NODE_ENV === 'production' ? basePath : ''}/assets/images/flags/spain_flag.jpg`}
                        alt="es"
                      />
                    </span>
                    Español
                  </div>
                </a>
              </li>
              <li>
                <a
                  className={`ti-dropdown-item flex items-center ${lang === "en" ? "bg-primary/10" : ""}`}
                  href="#!"
                  onClick={(e: React.MouseEvent) => { e.preventDefault(); setLanguage("en"); }}
                >
                  <div className="flex items-center">
                    <span className="avatar avatar-rounded avatar-xs leading-none me-2">
                      <img
                        src={`${process.env.NODE_ENV === 'production' ? basePath : ''}/assets/images/flags/us_flag.jpg`}
                        alt="en"
                      />
                    </span>
                    English
                  </div>
                </a>
              </li>
            </SpkDropdown>
            {/* <!-- End::header-element --> */}

            {/* <!-- Start::header-element --> */}
            {/* <!-- light and dark theme --> */}
            <li className="header-element header-theme-mode hidden !items-center sm:block md:!px-[0.5rem] px-2" onClick={() => toggledark()}>
              <Link scroll={false}
                aria-label="anchor"
                className="hs-dark-mode-active:hidden flex hs-dark-mode group flex-shrink-0 justify-center items-center gap-2  rounded-full font-medium transition-all text-xs dark:bg-bgdark dark:hover:bg-black/20 text-textmuted dark:text-textmuted/50 dark:hover:text-white dark:focus:ring-white/10 dark:focus:ring-offset-white/10"
                href="#!"
                data-hs-theme-click-value="dark"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-6 h-6 header-link-icon"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z"
                  />
                </svg>
              </Link>
              <Link scroll={false}
                aria-label="anchor"
                className="hs-dark-mode-active:flex hidden hs-dark-mode group flex-shrink-0 justify-center items-center gap-2  rounded-full font-medium text-defaulttextcolor  transition-all text-xs dark:bg-bodybg dark:bg-bgdark dark:hover:bg-black/20  dark:hover:text-white dark:focus:ring-white/10 dark:focus:ring-offset-white/10"
                href="#!"
                data-hs-theme-click-value="light"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-6 h-6 header-link-icon"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M12 3v2.25m6.364.386-1.591 1.591M21 12h-2.25m-.386 6.364-1.591-1.591M12 18.75V21m-4.773-4.227-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z"
                  />
                </svg>
              </Link>
            </li>
            {/* <!-- End light and dark theme --> */}

            {/* <!-- End::header-element --> */}

            {/* <!-- Start::header-element --> */}
            <li className="header-element header-fullscreen" onClick={toggleFullScreen}>
              {/* <!-- Start::header-link --> */}
              <Link scroll={false}
                // onclick="openFullscreen();"
                href="#!"
                className="header-link"
                aria-label="anchor"
              >
                {fullScreen ? (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="w-6 h-6 full-screen-close header-link-icon"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M9 9V4.5M9 9H4.5M9 9 3.75 3.75M9 15v4.5M9 15H4.5M9 15l-5.25 5.25M15 9h4.5M15 9V4.5M15 9l5.25-5.25M15 15h4.5M15 15v4.5m0-4.5 5.25 5.25"
                    />
                  </svg>

                ) : (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="w-6 h-6 full-screen-open header-link-icon"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M3.75 3.75v4.5m0-4.5h4.5m-4.5 0L9 9M3.75 20.25v-4.5m0 4.5h4.5m-4.5 0L9 15M20.25 3.75h-4.5m4.5 0v4.5m0-4.5L15 9m5.25 11.25h-4.5m4.5 0v-4.5m0 4.5L15 15"
                    />
                  </svg>
                )}
              </Link>
              {/* <!-- End::header-link --> */}
            </li>
            {/* <!-- End::header-element --> */}

            {/* <!-- Start::header-element --> */}
            <SpkDropdown Customclass="header-element" Linktag={true} Navigate='#!' Linkclass='header-link hs-dropdown-toggle ti-dropdown-toggle' Id="mainHeaderProfile" Imagetag={true}
              Imageclass='flex items-center avatar avatar-sm mb-0' Image={`${process.env.NODE_ENV === 'production' ? basePath : ''}/assets/images/faces/15.jpg`}
              Custommenuclass='main-header-dropdown  pt-0 overflow-hidden header-profile-dropdown ' Menulabel='mainHeaderProfile'>

              {/* <!-- End::header-link|dropdown-toggle --> */}

              <li>
                <div className="ti-dropdown-item text-center border-b border-defaultborder dark:border-defaultborder/10 block">
                  <span>Mr.Henry</span>
                  <span className="block text-xs text-textmuted dark:text-textmuted/50">
                    UI/UX Designer
                  </span>
                </div>
              </li>
              <li>
                <Link scroll={false}
                  className="ti-dropdown-item flex items-center"
                  href="#!"
                >
                  <i className="fe fe-user p-1 rounded-full bg-primary/10 text-primary me-2 text-[1rem]"></i>
                  Profile
                </Link>
              </li>
              <li>
                <Link scroll={false}
                  className="ti-dropdown-item flex items-center"
                  href="#!"
                >
                  <i className="fe fe-mail p-1 rounded-full bg-primary/10 text-primary me-2 text-[1rem]"></i>
                  Mail Inbox
                </Link>
              </li>
              <li>
                <Link scroll={false}
                  className="ti-dropdown-item flex items-center"
                  href="#!"
                >
                  <i className="fe fe-database p-1 rounded-full bg-primary/10 text-primary klist me-2 text-[1rem]"></i>
                  File Manager
                  <SpkBadge variant="primarytint1color" customClass="text-white ms-auto text-[0.5625rem]">
                    2
                  </SpkBadge>
                </Link>
              </li>
              <li>
                <Link scroll={false}
                  className="ti-dropdown-item flex items-center"
                  href="#!"
                >
                  <i className="fe fe-settings p-1 rounded-full bg-primary/10 text-primary ings me-2 text-[1rem]"></i>
                  Settings
                </Link>
              </li>
              <li className="border-t border-defaultborder dark:border-defaultborder/10 bg-light">
                <Link scroll={false}
                  className="ti-dropdown-item flex items-center"
                  href="#!"
                >
                  <i className="fe fe-help-circle p-1 rounded-full bg-primary/10 text-primary set me-2 text-[1rem]"></i>
                  Help
                </Link>
              </li>
              <li>
                <Link scroll={false}
                  className="ti-dropdown-item flex items-center"
                  href="#!"
                >
                  <i className="fe fe-lock p-1 rounded-full bg-primary/10 text-primary ut me-2 text-[1rem]"></i>
                  Log Out
                </Link>
              </li>
            </SpkDropdown>
            {/* <!-- End::header-element --> */}

            {/* <!-- Start::header-element --> */}
            <li className="header-element">
              {/* <!-- Start::header-link|switcher-icon --> */}
              <a
                href="#!"
                className="header-link switcher-icon"
                data-hs-overlay="#hs-overlay-switcher"
                aria-label="anchor"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-6 h-6 header-link-icon"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.325.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 0 1 1.37.49l1.296 2.247a1.125 1.125 0 0 1-.26 1.431l-1.003.827c-.293.241-.438.613-.43.992a7.723 7.723 0 0 1 0 .255c-.008.378.137.75.43.991l1.004.827c.424.35.534.955.26 1.43l-1.298 2.247a1.125 1.125 0 0 1-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.47 6.47 0 0 1-.22.128c-.331.183-.581.495-.644.869l-.213 1.281c-.09.543-.56.94-1.11.94h-2.594c-.55 0-1.019-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 0 1-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 0 1-1.369-.49l-1.297-2.247a1.125 1.125 0 0 1 .26-1.431l1.004-.827c.292-.24.437-.613.43-.991a6.932 6.932 0 0 1 0-.255c.007-.38-.138-.751-.43-.992l-1.004-.827a1.125 1.125 0 0 1-.26-1.43l1.297-2.247a1.125 1.125 0 0 1 1.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.086.22-.128.332-.183.582-.495.644-.869l.214-1.28Z"
                  />
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"
                  />
                </svg>
              </a>
              {/* <!-- End::header-link|switcher-icon --> */}
            </li>
            {/* <!-- End::header-element --> */}
          </ul>
          {/* <!-- End::header-content-right --> */}
        </div>
        {/* <!-- End::main-header-container --> */}

      </header>
      <div className="hs-overlay ti-modal hidden" id="header-responsive-search" tabIndex={-1} aria-labelledby="header-responsive-search" aria-hidden="true">
        <div className="ti-modal-box">
          <div className="ti-modal-dialog">
            <div className="ti-modal-content">
              <div className="ti-modal-body">
                <div className="input-group">
                  <input type="text" className="form-control border-end-0 !border-s" placeholder="Search Anything ..."
                    aria-label="Search Anything ..." aria-describedby="button-addon2" />
                  <SpkButton variant="primary" customClass="ti-btn !m-0" buttontype="button"
                    Id="button-addon2"><i className="bi bi-search"></i></SpkButton>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

    </Fragment>
  );
};

const mapStateToProps = (state: any) => ({
  local_varaiable: state,
});
export default connect(mapStateToProps, { ThemeChanger })(Header);
