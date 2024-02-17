#ifndef __invocationlistener__h__
#define __invocationlistener__h__
#include <frida-gum.h>
typedef struct _GumInvocationListenerProxyClass GumInvocationListenerProxyClass;
struct _GumInvocationListenerProxyClass
{
    GObjectClass parent_class;
};

struct _GumInvocationListenerProxy
{
    GObject parent;
    char Id[120];
};
typedef struct _GumInvocationListenerProxy GumInvocationListenerProxy;
static GType gum_invocation_listener_proxy_get_type();
static void gum_invocation_listener_proxy_iface_init(gpointer g_iface, gpointer iface_data);
void* CreateListener(char* id);
typedef void (*GumInvocationListenerProxyCallbackOnEnter)(void*,void*);
typedef void (*GumInvocationListenerProxyCallbackOnLeave)(void*,void*);
void SetGumInvocationProxyCallback(void* cbOnEnter,void* cbOnLeave);
#endif
