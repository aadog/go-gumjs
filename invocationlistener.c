#include "invocationlistener.h"
#include <stdio.h>
GumInvocationListenerProxyCallbackOnEnter ListenerProxyCbOnEnter;
GumInvocationListenerProxyCallbackOnLeave ListenerProxyCbOnLeave;

void SetGumInvocationProxyCallback(void* cbOnEnter,void* cbOnLeave){
    ListenerProxyCbOnEnter=cbOnEnter;
    ListenerProxyCbOnLeave=cbOnLeave;
}

G_DEFINE_TYPE_EXTENDED(GumInvocationListenerProxy,
                          gum_invocation_listener_proxy,
                          G_TYPE_OBJECT,
                          0,
                          G_IMPLEMENT_INTERFACE(GUM_TYPE_INVOCATION_LISTENER,
                              gum_invocation_listener_proxy_iface_init))

static void gum_invocation_listener_proxy_init(GumInvocationListenerProxy * self){

}

static void gum_invocation_listener_proxy_finalize(GObject * obj)
{
    G_OBJECT_CLASS(gum_invocation_listener_proxy_parent_class)->finalize(obj);
}

static void gum_invocation_listener_proxy_class_init(GumInvocationListenerProxyClass * klass)
{
    G_OBJECT_CLASS(klass)->finalize = gum_invocation_listener_proxy_finalize;
}

static void gum_invocation_listener_proxy_on_enter(GumInvocationListener * listener,GumInvocationContext * context)
{
    GumInvocationListenerProxy* mListener=(GumInvocationListenerProxy*)listener;
    ListenerProxyCbOnEnter(mListener,context);
}

static void gum_invocation_listener_proxy_on_leave(GumInvocationListener * listener,GumInvocationContext* context)
{
    GumInvocationListenerProxy* mListener=(GumInvocationListenerProxy*)listener;
    ListenerProxyCbOnLeave(mListener,context);
}

static void gum_invocation_listener_proxy_iface_init(gpointer g_iface,gpointer iface_data)
{
    GumInvocationListenerInterface * iface =(GumInvocationListenerInterface*)(g_iface);
    iface->on_enter = gum_invocation_listener_proxy_on_enter;
    iface->on_leave = gum_invocation_listener_proxy_on_leave;
}
void* CreateListener(char* id){
    GumInvocationListener* listener=g_object_new(gum_invocation_listener_proxy_get_type(), NULL);
    GumInvocationListenerProxy* mListener=(GumInvocationListenerProxy*)listener;
    strcpy(mListener->Id,id);
    return listener;
}
