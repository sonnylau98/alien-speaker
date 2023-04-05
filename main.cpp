#include<fstream>
#include<iostream>
#include <cstdlib>

using std::string;
using std::cout;
using std::ofstream;
using std::ifstream;
using std::cin;
using std::endl;

unsigned int key=0;
string EN(string en,unsigned int k);
string DE(string de,unsigned int k);

char dec2hexChar(short int n)
{
    if ( 0 <= n && n <= 9 )
    {
        return char( short('0') + n );
    }
	else if ( 10 <= n && n <= 15 )
	{
        return char( short('A') + n - 10 );
    }
	else
	{
        return char(0);
    }
}

short int hexChar2dec(char c)
{
    if ( '0'<=c && c<='9' )
	{
        return short(c-'0');
    }
	else if ( 'a'<=c && c<='f' )
	{
        return ( short(c-'a') + 10 );
    }
	else if ( 'A'<=c && c<='F' )
	{
        return ( short(c-'A') + 10 );
    }
	else
    {
        return -1;
    }
}

string escapeURL(const string &URL)
{
    string result = "";
    for ( unsigned int i=0; i<URL.size(); i++ )
	{
        char c = URL[i];
        if (( '0'<=c && c<='9' ) ||( 'a'<=c && c<='z' ) ||( 'A'<=c && c<='Z' ) ||c=='/' || c=='.')
	    {
            result += c;
        }
	    else
	    {
            int j = (short int)c;
            if ( j < 0 )
		    {
                j += 256;
            }
            int i1, i0;
            i1 = j / 16;
            i0 = j - i1*16;
            result += '%';
            result += dec2hexChar(i1);
            result += dec2hexChar(i0);
        }
    }
    return result;
}

string deescapeURL(const string &URL)
{
    string result = "";
    for ( unsigned int i=0; i<URL.size(); i++ )
	{
        char c = URL[i];
        if ( c != '%' )
		{
            result += c;
        }
		else
		{
            char c1 = URL[++i];
            char c0 = URL[++i];
            int num = 0;
            num += hexChar2dec(c1) * 16 + hexChar2dec(c0);
            result += char(num);
        }
    }
    return result;
}

int main()
{
	void fun1();
	void fun2();
	cout<<"*欢迎使用钩码翻译器Gamma版 :)*\n";
	cout<<"\n请输入密钥（非负整数）：";
	if(scanf("%u",&key)==0)
    {
        cout<<"\n输入有误:(\n";
        system("pause");
        exit(1);
    }
	int num;
	while(1)
	{
		cout<<"\n输入明文请按【1】，翻译密文请按【2】,更换密钥请按【3】，退出程序请按【4】。\n请输入选项：";
		if(scanf("%d",&num)==0)
		{
			cout<<"\n输入有误:(\n";
			system("pause");
			exit(1);
		 }
		switch(num)
		{
			case 1 : fun1();
			    break;
			case 2 : fun2();
			    break;
			case 3 :
			{
			    cout<<"\n请输入密钥（非负整数）：";
	            if(scanf("%u",&key)==0)
		        {
                    cout<<"\n输入有误:(\n";
                    system("pause");
			        exit(1);
		        }
			}
			    break;
			case 4 :
			{
				cout<<"\n感谢您的使用:),作者将继续改进此程序，从而给您带来更好的体验。\n\n";
				system("pause");
				exit(0);
			}
			    break;
			default : cout<<"\n警告：请按照指示输入选项。\n";
		 }
	}

	return 0;
 }

void fun1()
{
	string Text="";
	string cryptograph="";
	string mi="";
	cout<<"\n请输入要【加密】的明文：\n";
    cin.ignore();
	getline(cin,Text);
	cryptograph=escapeURL(Text);
	mi=EN(cryptograph,key);
    ofstream outfile;
    outfile.open("密文.txt");
    outfile << mi << endl;
    cout<<"\nOK了！加密后的密文已位于本目录下名为“密文”的文件里。\n";
    outfile.close();
}

void fun2()
{
	string Text="";
	string cryptograph="";
	string ming="";
	cout<<"\n请将需要翻译的【密文】输入到本目录下名为“密文”的文件里。\n确认后";
	system("pause");
	ifstream infile;
    infile.open("密文.txt");
    infile >>cryptograph;
    ming=DE(cryptograph,key);
    Text=deescapeURL(ming);
	cout<<"\n解密后的明文是：\n"<<Text<<endl;
    infile.close();
}

string EN(string en,unsigned int k)                //对URL加密
{
	string E;
	int i,j;
	for(i=0;i<en.size();)
	{
		for(j=0;(i<en.size())&&(j<32);i++,j++)
		{
			if(i==0)
			{
				j=k%32;
			}
			E.push_back(en[i]+j);
		}
	}
	return E;
}

string DE(string de,unsigned int k)
{
	string D;
	int i,j;
	for(i=0;i<de.size();)
	{
		for(j=0;(i<de.size())&&(j<32);i++,j++)
		{
			if(i==0)
			{
				j=k%32;
			}
			D.push_back(de[i]-j);
		}
	}
	return D;
}
