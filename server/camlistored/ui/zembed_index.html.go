// THIS FILE IS AUTO-GENERATED FROM index.html
// DO NOT EDIT.

package newui

import "time"

import "camlistore.org/pkg/fileembed"

func init() {
	Files.Add("index.html", 1213, time.Unix(0, 1370942742232957700), fileembed.String("<!doctype html>\n"+
		"<html>\n"+
		"  <head>\n"+
		"    <script src=\"closure/goog/base.js\"></script>\n"+
		"    <script src=\"./deps.js\"></script>\n"+
		"    <script src=\"?camli.mode=config&var=CAMLISTORE_CONFIG\"></script>\n"+
		"\n"+
		"    <!-- Begin non-Closure cheating; but depended on by server_connection.js -->\n"+
		"    <script type=\"text/javascript\" src=\"base64.js\"></script>\n"+
		"    <script type=\"text/javascript\" src=\"Crypto.js\"></script>\n"+
		"    <script type=\"text/javascript\" src=\"SHA1.js\"></script>\n"+
		"    <!-- End non-Closure cheating -->\n"+
		"\n"+
		"    <script>\n"+
		"      goog.require('camlistore.IndexPage');\n"+
		"    </script>\n"+
		"    <link rel=\"stylesheet\" href=\"blob_item.css\" type=\"text/css\">\n"+
		"    <link rel=\"stylesheet\" href=\"blob_item_container.css\" type=\"text/css\">\n"+
		"    <link rel=\"stylesheet\" href=\"create_item.css\" type=\"text/css\">\n"+
		"    <link rel=\"stylesheet\" href=\"index.css\" type=\"text/css\">\n"+
		"    <link rel=\"stylesheet\" href=\"toolbar.css\" type=\"text/css\">\n"+
		"    <link rel=\"stylesheet\" href=\"closure/goog/css/common.css\" type=\"text/css\">\n"+
		"    <link rel=\"stylesheet\" href=\"closure/goog/css/toolbar.css\" type=\"text/css\">\n"+
		"  </head>\n"+
		"  <body>\n"+
		"    <script>\n"+
		"      var page = new camlistore.IndexPage(CAMLISTORE_CONFIG);\n"+
		"      page.decorate(document.body);\n"+
		"    </script>\n"+
		"  </body>\n"+
		"</html>\n"+
		""))
}
